package logger

import "app/interface/config"

type LoggerHandler interface {
	Set(config.ConfigHandler) error
	Info(argv ...interface{})
	Debug(argv ...interface{})
	Error(argv ...interface{})
}
