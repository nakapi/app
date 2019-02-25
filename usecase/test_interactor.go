package usecase

import (
	"app/infrastructure/database/repository"
	"app/interface/presenter"
	"fmt"
)

type TestInteractor struct {
	TestRepository repository.TestRepository
	TestPresenter  presenter.TestPresenter
}

func NewTestInteractor(repository repository.TestRepository, presenter presenter.TestPresenter) *TestInteractor {
	interactor := new(TestInteractor)
	interactor.TestRepository = repository
	interactor.TestPresenter = presenter
	return interactor
}

func (interactor TestInteractor) Handle() {
	tests, err := interactor.TestRepository.FindAll()
	if err != nil {
		interactor.TestPresenter.TestOutputData.Error = fmt.Errorf("Find ALl Failed %s", err.Error())
	}
	interactor.TestPresenter.TestOutputData.Id = tests[0].Id
	interactor.TestPresenter.TestOutputData.Name = tests[0].Name
	interactor.TestPresenter.Complete()
	return
}
