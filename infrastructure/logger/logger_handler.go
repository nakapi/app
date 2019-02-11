package logger

import (
	"app/interface/config"
	"app/interface/logger"
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type LoggerHandler struct {
	Logger *zap.Logger
}

func NewLoggerHandler() logger.LoggerHandler {
	handler := new(LoggerHandler)
	return handler
}

func (handler *LoggerHandler) Set(config config.ConfigHandler) error {
	var appConfig zap.Config
	err := json.Unmarshal(config.GetLogger(), &appConfig)
	if err != nil {
		fmt.Println("bbb")
		return fmt.Errorf("Logger Config is incorrect")
	}
	handler.Logger, _ = appConfig.Build()
	return nil
}

func (handler *LoggerHandler) Info() {
	handler.Logger.Info("Hello Zap", zap.String("key", "value"), zap.Time("now", time.Now()))
}
