package logger_test

import (
	"io/fs"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/mclargo/go-framework/cmd/conf"
	"github.com/mclargo/go-framework/internal/logger"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type testLogSuite struct {
	suite.Suite
	config *conf.Config
}

func TestLogSuite(t *testing.T) {
	suite.Run(t, new(testLogSuite))
}

func (s *testLogSuite) SetupSuite() {
	s.config = &conf.Config{
		App: conf.AppConfig{
			Verbose: new(bool),
			Debug:   new(bool),
			Port:    ":3000",
		},
		Log: conf.LogConfig{
			Path:     "../../tmp/test",
			Filename: "test.log",
			Debug:    new(bool),
		},
	}
}

func (s *testLogSuite) TearDownTest() {
	// remove log path
	os.RemoveAll(s.config.Log.Path)
	// reset monkey
	monkey.UnpatchAll()
}

func (s *testLogSuite) TestInitLogger() {
	tt := []struct {
		name  string
		debug bool
	}{
		{
			name:  "TestInitLogger_DebugTrue",
			debug: true,
		},
		{
			name:  "TestInitLogger_DebugFalse",
			debug: false,
		},
	}

	for _, tc := range tt {
		s.T().Run(tc.name, func(_ *testing.T) {
			// Arrange
			s.config.Log.Debug = &tc.debug

			// Act
			log, err := logger.InitLogger(*s.config)

			// Assert
			s.Require().NoError(err)
			s.Require().NotNil(log)
			s.Require().Equal(log.Core().Enabled(zap.DebugLevel), tc.debug)

			// combine original logger with observer
			level := zap.InfoLevel
			if tc.debug {
				level = zap.DebugLevel
			}

			coreObserver, logs := observer.New(level)

			newCore := zapcore.NewTee(log.Core(), coreObserver)
			newLog := zap.New(newCore)

			// Check log levels
			newLog.Debug("Test")
			newLog.Info("Test")
			newLog.Warn("Test")
			newLog.Error("Test")

			// Assert
			if tc.debug {
				s.Require().Equal(4, logs.Len())
			} else {
				s.Require().Equal(3, logs.Len())
			}
		})
	}
}

func (s *testLogSuite) TestInitLogger_CreatePath_KO() {
	// Arrange
	monkey.Patch(os.Mkdir, func(_ string, _ fs.FileMode) error {
		return fs.ErrExist
	})

	// Act
	log, err := logger.InitLogger(*s.config)

	// Assert
	s.Require().Error(err)
	s.Require().Nil(log)
}

func (s *testLogSuite) TestInitLogger_OpenFile_KO() {
	// Arrange
	monkey.Patch(os.OpenFile, func(_ string, _ int, _ fs.FileMode) (*os.File, error) {
		return nil, fs.ErrNotExist
	})

	// Act
	log, err := logger.InitLogger(*s.config)

	// Assert
	s.Require().Error(err)
	s.Require().Nil(log)
}
