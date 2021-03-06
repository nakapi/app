package log

import (
	"app/interface/config"
	"app/interface/logger"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerHandler struct {
	Logger *zap.Logger
}

func NewLoggerHandler(config config.ConfigHandler) (logger.LoggerHandler, error) {
	handler := new(LoggerHandler)
	err := handler.Set(config)
	if err != nil {
		return nil, fmt.Errorf("Log Config Error", err.Error())
	}
	return handler, nil
}

func (handler *LoggerHandler) Set(config config.ConfigHandler) error {
	var appConfig zap.Config
	err := json.Unmarshal(config.GetLogger(), &appConfig)
	if err != nil {
		return fmt.Errorf("Logger Config is incorrect")
	}
	handler.Logger, _ = appConfig.Build()
	return nil
}

func (handler LoggerHandler) Info(argv ...interface{}) {
	firstArgv := GetFirstArgv(argv...)
	optionArgv := GetOptionArgv(argv...)
	handler.Logger.Info(firstArgv, optionArgv...)
}

func (handler LoggerHandler) Debug(argv ...interface{}) {
	firstArgv := GetFirstArgv(argv...)
	optionArgv := GetOptionArgv(argv...)
	handler.Logger.Debug(firstArgv, optionArgv...)
}

func (handler LoggerHandler) Error(argv ...interface{}) {
	firstArgv := GetFirstArgv(argv...)
	optionArgv := GetOptionArgv(argv...)
	handler.Logger.Error(firstArgv, optionArgv...)
}

func GetFirstArgv(argv ...interface{}) string {
	firstArgv, _ := argv[0].(string)
	return firstArgv
}

func GetOptionArgv(argv ...interface{}) []zapcore.Field {
	var optionsArgv []zapcore.Field
	for _, value := range argv[1:] {
		configValue, ok := value.(zapcore.Field)
		if ok {
			optionsArgv = append(optionsArgv, configValue)
		}
	}
	return optionsArgv
}
