package db

import (
	"github.com/mclargo/go-framework/cmd/conf"
	"go.uber.org/zap"
)

type InMemoryAdapter struct {
	config *conf.Config
	logger *zap.Logger
}

func NewInMemoryAdapter(c *conf.Config, l *zap.Logger) *InMemoryAdapter {
	return &InMemoryAdapter{
		config: c,
		logger: l,
	}
}

var _ DatabaseAdapter = (*InMemoryAdapter)(nil)

func (i *InMemoryAdapter) Connect() error {
	i.logger.Debug("connecting to in-memory database")

	return nil
}

func (i *InMemoryAdapter) Disconnect() error {
	i.logger.Debug("disconnecting from in-memory database")

	return nil
}
