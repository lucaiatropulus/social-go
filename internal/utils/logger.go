package utils

import (
	"go.uber.org/zap"
	"sync"
)

var (
	logger *zap.SugaredLogger
	once   sync.Once
)

func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		newLogger, _ := zap.NewProduction()
		logger = newLogger.Sugar()
	})
	return logger
}
