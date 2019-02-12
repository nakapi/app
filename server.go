package main

import (
	file "app/infrastructure/config/json"
	"app/infrastructure/database"
	"app/infrastructure/database/repository"
	"app/infrastructure/logger"
	"app/usecase"
	"time"

	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
)

const (
	noError           = iota // OK
	errorDBConnection        // NG
	errorConfig              // NG
)

func main() {
	client := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(client.Run(os.Args))
}

type CLI struct {
	outStream io.Writer
	errStream io.Writer
}

func (client *CLI) Run(args []string) int {
	// Config
	conf := file.NewConfigHandler()
	err := conf.Load()
	if err != nil {
		fmt.Println("Config Load Failed %s", err.Error())
		return errorConfig
	}
	// Log
	//	var appLogConfig zap.Config
	logger := logger.NewLoggerHandler()
	logger.Set(conf)
	logger.Info("Test", zap.String("key", "value"), zap.Time("now", time.Now()))
	logger.Error("Test")

	// Database
	sqlHandler, err := database.NewSqlHandler(conf)
	if err != nil {
		logger.Error("SqlHandler instance Failed. %s", err.Error())
		return errorDBConnection
	}

	interactor := usecase.TestInteractor{
		TestRepository: repository.TestRepository{
			SqlHandler: sqlHandler,
		},
	}
	tests, err := interactor.Tests()
	fmt.Println(err)
	fmt.Println(tests)

	/*
		res, err := sqlHandler.Query("select * from test")
		defer res.Close()
		if err != nil {
			logger.Error("SqlHandler Query Failed. %s", err.Error())
			return errorDBConnection
		}
		var id int
		var name string
		res.Next()
		if err = res.Scan(&id, &name); err != nil {
			fmt.Println(err.Error())
			return errorDBConnection
		}
		fmt.Println(id)
		fmt.Println(name)
	*/

	return noError
}
