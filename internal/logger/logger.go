package logger

import (
	"go.uber.org/zap"
)

func New(env string) *zap.Logger {
	var cfg zap.Config

	if env == "loacl" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
