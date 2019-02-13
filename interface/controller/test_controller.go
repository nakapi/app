package controller

import (
	"app/infrastructure/database/repository"
	"app/interface/database"
	"app/usecase"
	"context"
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
		},
	}
}

func (controller *TestController) Index(ctx *context.Context) {
	tests, err := controller.Interactor.Tests()
	if err != nil {
		return
	}
	*ctx = context.WithValue(*ctx, "id", tests[0].Id)
	*ctx = context.WithValue(*ctx, "name", tests[0].Name)
	return
}
