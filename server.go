// メインパッケージ
// main
package main

//  インポートパッケージ
import (
	"app/infrastructure/dicontainer"
	"app/interface/controller"
	"app/interface/logger"
	"fmt"

	"io"
	"os"
)

// コマンド結果リスト
const (
	noError           = iota // OK
	errorDBConnection        // NG
	errorConfig              // NG
)

// クライアントコマンドインタフェース構造
type CLI struct {
	// 標準出力先
	outStream io.Writer
	// エラー出力先
	errStream io.Writer
}

// メイン処理を行います
func main() {
	client := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(client.Run(os.Args))
}

// メイン処理の中で、実際の処理を担い結果を出力します。
func (client *CLI) Run(args []string) int {

	// DIContainer
	// TODO: DI Error Handling
	container := dicontainer.NewContainerHandler()

	// Logging begin
	logContainer, err := container.Resolve("logger")
	if err != nil {
		fmt.Println(err.Error())
		return errorConfig
	}
	loggerHandler := logContainer.(logger.LoggerHandler)
	loggerHandler.Info("App Begin")

	// Controller:Controller->UseCase(Interactor)->Repository(findAll)->Domain(Tests->Test) ===> Context Return
	controllerContainer, err := container.Resolve("testController")
	if err != nil {
		loggerHandler.Error("Controller DI Container Resolver Failed ", err.Error())
		return errorConfig
	}
	testController := controllerContainer.(*controller.TestController)
	testController.Index()

	// End
	loggerHandler.Info("App End")
	return noError
}
