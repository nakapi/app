// メインパッケージ
// main
package main

//  インポートパッケージ
import (
	"app/domain"
	"app/infrastructure/config/json"
	"app/infrastructure/database"
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

	// ElasticSearch
	config, err := json.NewConfigHandler()
	if err != nil {
		fmt.Println("Config Failed", err.Error())
	}
	elclient, err := database.NewSqlHandler(config)
	if err != nil {
		fmt.Println("Elastic Client Failed", err.Error())
	}
	result, err := elclient.Query("id", "2")
	if err != nil {
		fmt.Println("Elastic Query Failed", err.Error())
	}
	var testObject domain.Test
	result.Scan(&testObject)
	fmt.Println(testObject)

	return noError

	// DIContainer
	container, err := dicontainer.NewContainerHandler()
	if err != nil {
		fmt.Println(err.Error())
		return errorConfig
	}

	// Logging begin
	loggerContainer, err := container.Resolve("logger")
	if err != nil {
		fmt.Println(err.Error())
		return errorConfig
	}
	loggerHandler := loggerContainer.(logger.LoggerHandler)
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
