package main

import (
	file "app/infrastructure/config/json"
	"app/infrastructure/logger"
	"app/infrastructure/mysql"

	"fmt"
	"io"
	"os"
)

const (
	noError           = iota // OK
	errorDBConnection        // NG
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
	fmt.Println("main start")

	// Config
	conf := file.NewConfigHandler()
	err := conf.Load()
	if err != nil {
		fmt.Println(err.Error())
	}
	// Log
	//	var appLogConfig zap.Config
	logger := logger.NewLoggerHandler()
	logger.Set(conf)
	logger.Info()

	// Database
	sqlHandler, err := mysql.NewSqlHandler(conf)
	if err != nil {
		fmt.Println(err.Error())
		return errorDBConnection
	}
	res, err := sqlHandler.Query("select * from test")
	defer res.Close()
	if err != nil {
		fmt.Println(err.Error())
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

	fmt.Println("main end")
	return noError
}
