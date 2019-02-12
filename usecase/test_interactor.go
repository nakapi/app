package usecase

import (
	"app/domain"
	"app/infrastructure/database/repository"
)

type TestInteractor struct {
	TestRepository repository.TestRepository
}

func (interactor *TestInteractor) Add(test domain.Test) (insertTest domain.Test, err error) {
	identifier, err := interactor.TestRepository.Store(test)
	if err != nil {
		return
	}
	test, err = interactor.TestRepository.FindById(identifier)
	return
}

func (interactor *TestInteractor) Tests() (tests domain.Tests, err error) {
	tests, err = interactor.TestRepository.FindAll()
	return
}

func (interactor *TestInteractor) TestById(identifier int) (test domain.Test, err error) {
	test, err = interactor.TestRepository.FindById(identifier)
	return
}
