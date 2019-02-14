package usecase

import (
	"app/domain"
	"app/infrastructure/database/repository"
	"app/interface/presenter"
	"fmt"
)

type TestInteractor struct {
	TestRepository repository.TestRepository
	TestPresenter  presenter.TestPresenter
}

func (interactor *TestInteractor) Add(test domain.Test) (insertTest domain.Test, err error) {
	identifier, err := interactor.TestRepository.Store(test)
	if err != nil {
		return
	}
	test, err = interactor.TestRepository.FindById(identifier)
	return
}

func (interactor *TestInteractor) Tests() {
	tests, err := interactor.TestRepository.FindAll()
	if err != nil {
		interactor.TestPresenter.TestOutputData.Error = fmt.Errorf("Find ALl Failed %s", err.Error())
	}
	interactor.TestPresenter.TestOutputData.Id = tests[0].Id
	interactor.TestPresenter.TestOutputData.Name = tests[0].Name
	interactor.TestPresenter.Complete()
	return
}

func (interactor *TestInteractor) TestById(identifier int) (test domain.Test, err error) {
	test, err = interactor.TestRepository.FindById(identifier)
	return
}
