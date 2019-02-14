package presenter

import (
	"app/interface/presenter/dto"
	"fmt"
)

type TestPresenter struct {
	TestOutputData dto.TestOuputData
}

func (presenter *TestPresenter) Complete() {
	if presenter.TestOutputData.Error != nil {
		fmt.Println(presenter.TestOutputData.Error.Error())
		return
	}
	fmt.Println("ID:", presenter.TestOutputData.Id)
	fmt.Println("NAME:", presenter.TestOutputData.Name)
}
