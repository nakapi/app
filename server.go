package main

import (
	file "app/infrastructure/config/json"
	"app/infrastructure/database"
	"app/infrastructure/logger"
	"app/interface/controller"
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
	logger := logger.NewLoggerHandler()
	logger.Set(conf)
	logger.Info("App Start", zap.String("key", "value"), zap.Time("now", time.Now()))

	// Database
	sqlHandler, err := database.NewSqlHandler(conf)
	if err != nil {
		logger.Error("SqlHandler instance Failed. %s", err.Error())
		return errorDBConnection
	}

	// Controller->UseCase(Interactor)->Repository(findAll)->Domain(Tests->Test) ===> Context Return
	controller := controller.NewTestController(sqlHandler)
	controller.Index()

	// End
	logger.Info("App End", zap.String("key", "value"), zap.Time("now", time.Now()))
	return noError
}
