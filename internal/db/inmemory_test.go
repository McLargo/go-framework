package db_test

import (
	"testing"

	"github.com/mclargo/go-framework/cmd/conf"
	"github.com/mclargo/go-framework/internal/db"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testInMemorySuite struct {
	suite.Suite
	config       *conf.Config
	logger       *zap.Logger
	logsObserver *observer.ObservedLogs
	adapter      *db.InMemoryAdapter
}

func TestInMemorySuite(t *testing.T) {
	suite.Run(t, new(testInMemorySuite))
}

func (s *testInMemorySuite) SetupSuite() {
	// Arrange con
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
		Storage: conf.StorageConfig{
			Type: "memory",
		},
	}

	// Arrange logger and observer
	coreObserver, logs := observer.New(zap.DebugLevel)
	s.logger = zap.New(coreObserver)
	s.logsObserver = logs

	s.adapter = db.NewInMemoryAdapter(s.config, s.logger)
}

func (s *testInMemorySuite) TearDownSuite() {
	err := s.adapter.Disconnect()
	s.Require().NoError(err)

	msg := "disconnecting from in-memory database"
	s.Require().Equal(1, s.logsObserver.FilterMessage(msg).Len())
}

func (s *testInMemorySuite) TestConnect() {
	// Act
	err := s.adapter.Connect()

	// Assert
	s.NotNil(s.adapter)
	s.Require().NoError(err)

	msg := "connecting to in-memory database"
	s.Require().Equal(1, s.logsObserver.FilterMessage(msg).Len())
}
