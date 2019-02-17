package controller

import (
	"app/infrastructure/database/repository"
	"app/interface/database"
	"app/interface/presenter"
	"app/usecase"
)

type TestController struct {
	Interactor usecase.TestInteractor
}

func NewTestController(sqlHandler database.SqlHandler) *TestController {
	return &TestController{
		Interactor: usecase.TestInteractor{
			TestRepository: repository.TestRepository{
				SqlHandler: sqlHandler,
			},
			TestPresenter: presenter.TestPresenter{},
		},
	}
}

func (controller TestController) Index() {
	controller.Interactor.Tests()
	return
}
