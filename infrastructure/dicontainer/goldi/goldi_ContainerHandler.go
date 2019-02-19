package goldi

import (
	"app/infrastructure/config/json"
	"app/infrastructure/database"
	"app/infrastructure/log"
	"app/interface/controller"
	"app/interface/di"

	"github.com/fgrosse/goldi"
)

type ContainerHandler struct {
	Container *goldi.Container
}

func NewContainerHandler() (di.DiContainer, error) {
	container := new(ContainerHandler)
	// DIContailer
	registry := goldi.NewTypeRegistry()
	config := map[string]interface{}{}
	container.Container = goldi.NewContainer(registry, config)
	// Config
	container.Container.Register("config", goldi.NewType(json.NewConfigHandler))
	// LOG
	container.Container.RegisterType("logger", log.NewLoggerHandler, ("@config"))
	// DB
	container.Container.RegisterType("dbHandler", database.NewSqlHandler, ("@config"))
	// Controller:Controller->UseCase(Interactor)->Repository(findAll)->Domain(Tests->Test) ===> Context Return
	container.Container.RegisterType("testController", controller.NewTestController, ("@dbHandler"))

	return container, nil
}

func (container *ContainerHandler) Register(containerBuilder ...interface{}) error {
	return nil
}

func (container *ContainerHandler) Resolve(containerName interface{}) (interface{}, error) {
	name := containerName.(string)
	return container.Container.Get(name)
}
