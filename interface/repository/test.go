package repository

import "app/domain"

type TestRepository interface {
	Store(domain.Test) (int, error)
	FindById(int) (domain.Test, error)
	FindAll() (domain.Tests, error)
}
