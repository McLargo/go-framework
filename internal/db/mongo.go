package db

import (
	"context"
	"fmt"
	"time"

	"github.com/mclargo/go-framework/cmd/conf"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.uber.org/zap"
)

type MongoAdapter struct {
	config         *conf.Config
	logger         *zap.Logger
	client         *mongo.Client
	secondsTimeout time.Duration
}

var _ DatabaseAdapter = (*MongoAdapter)(nil)

func NewMongoAdapter(c *conf.Config, l *zap.Logger) *MongoAdapter {
	return &MongoAdapter{
		config:         c,
		logger:         l,
		client:         nil,
		secondsTimeout: time.Duration(c.Storage.Timeout) * time.Second,
	}
}

// Connect establishes a connection with the mongo db defined in the configuration file.
func (m *MongoAdapter) Connect() error {
	m.logger.Debug("connecting to mongo database")
	ctx, cancel := context.WithTimeout(context.Background(), m.secondsTimeout)

	defer cancel()

	mc, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(m.config.Storage.MongoURI),
		options.Client().SetTimeout(m.secondsTimeout),
	)

	if err != nil {
		return fmt.Errorf("failed to create mongo db client for %w", err)
	}

	// send ping to database to confirm a successful connection
	m.logger.Debug("sending ping to mongo database")

	err = mc.Ping(context.Background(), readpref.Primary())

	if err != nil {
		m.logger.Error("MongoDB ping failed")
		return fmt.Errorf("failed to ping database: %w", err)
	}

	m.client = mc
	m.logger.Debug("connected to mongo database")

	return nil
}

func (m *MongoAdapter) Disconnect() error {
	m.logger.Debug("disconnecting from mongo database")

	if err := m.client.Disconnect(context.Background()); err != nil {
		return fmt.Errorf("failed to disconnect from mongo db: %w", err)
	}

	m.logger.Debug("disconnected from mongo database")

	return nil
}
