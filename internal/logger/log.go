package logger

import (
	"fmt"
	"os"

	"github.com/mclargo/go-framework/cmd/conf"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(cfg conf.Config) (*zap.Logger, error) {
	pe := zap.NewProductionEncoderConfig()

	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(pe)
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	level := zap.InfoLevel
	if *cfg.Log.Debug {
		level = zap.DebugLevel
	}

	pathToLog := fmt.Sprintf("%s/%s", cfg.Log.Path, cfg.Log.Filename)

	// create the log folder if it does not exist
	if _, err := os.Stat(pathToLog); os.IsNotExist(err) {
		os.Mkdir(cfg.Log.Path, os.ModePerm)
	}

	f, err := os.OpenFile(pathToLog, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("cannot open log file: %w", err)
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	logger := zap.New(core, zap.AddCaller())
	return logger, nil
}
