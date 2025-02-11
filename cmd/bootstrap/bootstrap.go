package bootstrap

import (
	"github.com/mclargo/go-framework/cmd/conf"
	"github.com/mclargo/go-framework/internal/logger"

	"go.uber.org/zap"
)

type Bootstrap struct {
	Config *conf.Config
	Logger *zap.Logger
}

func newBootstrap() (Bootstrap, error) {
	cfg, err := conf.InitConfig()
	if err != nil {
		return Bootstrap{}, err
	}

	log, err := logger.InitLogger(*cfg)
	if err != nil {
		return Bootstrap{}, err
	}

	if *cfg.App.Verbose {
		cfg.Print(log)
	}

	return Bootstrap{
		Config: cfg,
		Logger: log,
	}, nil
}
