package presenter

import (
	"app/interface/presenter/dto"
)

type ITestPresenter interface {
	Complete(dto.TestOutputData)
}
