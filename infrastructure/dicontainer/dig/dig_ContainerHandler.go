package dicontainer

import (
	"app/infrastructure/config/json"
	"app/infrastructure/database"
	"app/infrastructure/log"
	"app/interface/config"
	idatabase "app/interface/database"
	"app/interface/di"
	"app/interface/logger"
	"fmt"

	"go.uber.org/dig"
)

type ContainerHandler struct {
	Container *dig.Container
}

func NewContainerHandler() (di.DiContainer, error) {
	container := new(ContainerHandler)
	container.Container = dig.New()

	// ConfigHandler
	err := container.Container.Provide(func() (config.ConfigHandler, error) {
		return json.NewConfigHandler()
	}, dig.Name("config"))
	if err != nil {
		return nil, fmt.Errorf("Container Provide ConfigHandler Failed", err.Error())
	}

	// LogHandler
	err = container.Container.Provide(func(config config.ConfigHandler) (logger.LoggerHandler, error) {
		return log.NewLoggerHandler(config)
	}, dig.Name("logger"))
	if err != nil {
		return nil, fmt.Errorf("Container Provide LogHandler Failed", err.Error())
	}

	// DbHandler
	err = container.Container.Provide(func(config config.ConfigHandler) (idatabase.SqlHandler, error) {
		return database.NewSqlHandler(config)
	}, dig.Name("db"))
	if err != nil {
		return nil, fmt.Errorf("Container Provide DbHandler Failed", err.Error())
	}

	return container, nil
}

func (container *ContainerHandler) Register(containerType ...interface{}) error {
	return nil
}
func (container *ContainerHandler) Resolve(containerHandler interface{}) (interface{}, error) {
	return nil, nil
}
