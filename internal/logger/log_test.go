package logger_test

import (
	"os"
	"testing"

	"github.com/mclargo/go-framework/cmd/conf"
	"github.com/mclargo/go-framework/internal/logger"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type testLogSuite struct {
	suite.Suite
}

func TestLogSuite(t *testing.T) {
	suite.Run(t, new(testLogSuite))
}

var defaultToFalse = false

var cfg = conf.Config{
	App: conf.AppConfig{
		Verbose: &defaultToFalse,
		Debug:   &defaultToFalse,
		Port:    ":3000",
	},
	Log: conf.LogConfig{
		Path:     "../../tmp/test",
		Filename: "test.log",
		Debug:    &defaultToFalse,
	},
}

func (s *testLogSuite) TearDownTest() {
	os.RemoveAll(cfg.Log.Path)
}

func (s *testLogSuite) TestInitLogger() {
	tt := []struct {
		name  string
		cfg   conf.Config
		debug bool
	}{
		{
			name:  "TestInitLogger_DebugTrue",
			cfg:   cfg,
			debug: true,
		},
		{
			name:  "TestInitLogger_DebugFalse",
			cfg:   cfg,
			debug: false,
		},
	}

	for _, tc := range tt {
		s.T().Run(tc.name, func(_ *testing.T) {
			// Arrange
			tc.cfg.Log.Debug = &tc.debug

			// Arrange
			log, err := logger.InitLogger(tc.cfg)

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
