package presenter

import (
	"app/interface/presenter/dto"
	"bytes"
	"fmt"
	"io"
	"os/user"
	"path/filepath"
	"text/template"
)

type TestPresenter struct {
	TestOutputData dto.TestOuputData
}

func (presenter *TestPresenter) Complete() {
	if presenter.TestOutputData.Error != nil {
		fmt.Println(presenter.TestOutputData.Error.Error())
		return
	}
	user, err := user.Current()
	if err != nil {
		fmt.Println("Get User Failed ", err.Error())
		return
	}
	path := filepath.Join(user.HomeDir, "go", "src", "app", "infrastructure", "gui", "test.html")

	tmpl := template.Must(template.ParseFiles(path))
	buffer := new(bytes.Buffer)
	fw := io.Writer(buffer)
	dat := presenter.TestOutputData
	if err := tmpl.ExecuteTemplate(fw, "test", dat); err != nil {
		fmt.Println("Template Error ", err.Error())
		return
	}
	fmt.Println(string(buffer.Bytes()))

}
