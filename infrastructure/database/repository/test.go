package repository

import (
	"app/domain"
	"app/interface/database"
	"fmt"
)

type TestRepository struct {
	database.SqlHandler
}

func NewTestRepository(handler database.SqlHandler) *TestRepository {
	repository := new(TestRepository)
	repository.SqlHandler = handler
	return repository
}

func (repository *TestRepository) Store(test domain.Test) (id int, err error) {
	result, err := repository.Execute(
		"insert into test (id, name) values (?,?)", test.Id, test.Name,
	)
	if err != nil {
		return 0, fmt.Errorf("Test Insert Record Failed")
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Test Record Inserrt LastInsertId Failed")
	}
	id = int(insertId)
	return
}

func (repository *TestRepository) FindById(identifier int) (user domain.Test, err error) {
	row, err := repository.Query(
		"select id,name from test where id = ?", identifier,
	)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var name string
	if err = row.Scan(&id, &name); err != nil {
		return
	}
	user.Id = id
	user.Name = name
	return
}

func (repository *TestRepository) FindAll() (tests domain.Tests, err error) {
	rows, err := repository.Query(
		"select id,name from test",
	)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			continue
		}
		test := domain.Test{
			Id:   id,
			Name: name,
		}
		tests = append(tests, test)
	}
	return
}
