package db_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/mclargo/go-framework/cmd/conf"
	"github.com/mclargo/go-framework/internal/db"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

const (
	mongoImage   = "mongo:8.0"
	rootUsername = "root"
	rootPassword = "root"
)

type testMongoSuite struct {
	suite.Suite
	config       *conf.Config
	logger       *zap.Logger
	logsObserver *observer.ObservedLogs
	adapter      *db.MongoAdapter
}

func TestMongoSuite(t *testing.T) {
	suite.Run(t, new(testMongoSuite))
}

func (s *testMongoSuite) SetupSuite() {
	ctx := context.Background()
	// Init mongo in testcontainers
	req := testcontainers.ContainerRequest{
		Image:        mongoImage,
		ExposedPorts: []string{"27017/tcp", "27018/tcp"},
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": rootUsername,
			"MONGO_INITDB_ROOT_PASSWORD": rootPassword,
		},
		WaitingFor: wait.ForLog("MongoDB starting"),
	}

	mongoDBContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	s.Require().NoError(err)

	host, err := mongoDBContainer.Host(ctx)
	s.Require().NoError(err)
	p, err := mongoDBContainer.MappedPort(ctx, "27017/tcp")
	s.Require().NoError(err)

	port := p.Int()
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/", rootUsername, rootPassword, host, port)

	// Create temp folder
	dirTmp, err := os.MkdirTemp("", "log")
	s.Require().NoError(err)

	// Init config
	s.config = &conf.Config{
		App: conf.AppConfig{
			Verbose: new(bool),
			Debug:   new(bool),
			Port:    ":3000",
		},
		Log: conf.LogConfig{
			Path:     dirTmp,
			Filename: "test.log",
			Debug:    new(bool),
		},
		Storage: conf.StorageConfig{
			Type:     "mongodb",
			Timeout:  10,
			MongoURI: uri,
		},
	}

	// Init logger and observer
	coreObserver, logs := observer.New(zap.DebugLevel)
	s.logger = zap.New(coreObserver)
	s.logsObserver = logs

	s.adapter = db.NewMongoAdapter(s.config, s.logger)
}

func (s *testMongoSuite) TearDownSuite() {
	// Disconnect from MongoDB
	err := s.adapter.Disconnect()
	s.Require().NoError(err)

	msg := "disconnecting from mongo database"
	s.Require().Equal(1, s.logsObserver.FilterMessage(msg).Len())

	// Remove log path
	err = os.RemoveAll(s.config.Log.Path)
	s.Require().NoError(err)
}

func (s *testMongoSuite) TestConnect() {
	// Act
	err := s.adapter.Connect()

	// Assert
	s.NotNil(s.adapter)
	s.Require().NoError(err)

	msg := "connecting to mongo database"
	s.Require().Equal(1, s.logsObserver.FilterMessage(msg).Len())

	msg = "connected to mongo database"
	s.Require().Equal(1, s.logsObserver.FilterMessage(msg).Len())
}
