package dicontainer

import (
	"app/infrastructure/config/json"
	"app/infrastructure/database"
	"app/infrastructure/database/repository"
	"app/infrastructure/log"
	"app/interface/config"
	"app/interface/controller"
	idi "app/interface/di"
	"app/interface/presenter"
	"app/usecase"

	"github.com/sarulabs/di"
)

type ContainerHandler struct {
	Container di.Container
}

func NewContainerHandler() (idi.DiContainer, error) {
	containerHandler := new(ContainerHandler)
	builder, _ := di.NewBuilder()
	builder.Add([]di.Def{
		{
			Name: "config",
			Build: func(container di.Container) (interface{}, error) {
				return json.NewConfigHandler()
			},
		},
		{
			Name: "logger",
			Build: func(container di.Container) (interface{}, error) {
				return log.NewLoggerHandler(container.Get("config").(config.ConfigHandler))
			},
		},
		{
			Name: "database",
			Build: func(container di.Container) (interface{}, error) {
				return database.NewSqlHandler(container.Get("config").(config.ConfigHandler))
			},
		},
		{
			Name: "testController",
			Build: func(container di.Container) (interface{}, error) {
				return controller.NewTestController(container.Get("testInteractor").(usecase.ITestInteractor)), nil
			},
		},
		{
			Name: "testPresenter",
			Build: func(container di.Container) (interface{}, error) {
				return presenter.NewTestPresenter(), nil
			},
		},
		{
			Name: "testRepository",
			Build: func(container di.Container) (interface{}, error) {
				return repository.NewTestRepository(container.Get("database").(*database.SqlHandler)), nil
			},
		},
		{
			Name: "testInteractor",
			Build: func(container di.Container) (interface{}, error) {
				repository := container.Get("testRepository").(*repository.TestRepository)
				presenter := container.Get("testPresenter").(*presenter.TestPresenter)
				return usecase.NewTestInteractor(*repository, *presenter), nil
			},
		},
	}...,
	)
	containerHandler.Container = builder.Build()
	return containerHandler, nil
}

func (container *ContainerHandler) Register(registerArgv ...interface{}) error {
	return nil
}

func (container *ContainerHandler) Resolve(resolveArgv interface{}) (interface{}, error) {
	name := resolveArgv.(string)
	return container.Container.SafeGet(name)
}
