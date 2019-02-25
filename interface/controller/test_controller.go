package controller

import (
	"app/usecase"
)

type TestController struct {
	Interactor usecase.ITestInteractor
}

func NewTestController(interactor usecase.ITestInteractor) *TestController {
	controller := new(TestController)
	controller.Interactor = interactor
	return controller
}

func (controller TestController) Index() {
	controller.Interactor.Handle()
	return
}
