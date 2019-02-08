package main

import (
	"app/infrastructure/config"
	"app/infrastructure/mysql"
	"fmt"
)

func main() {
	fmt.Println("main start")
	conf := config.Load()
	sqlHandler := mysql.NewSqlHandler(conf)
	res, err := sqlHandler.Query("select * from test")
	defer res.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	var id int
	var name string
	res.Next()
	if err = res.Scan(&id, &name); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(id)
	fmt.Println(name)

	fmt.Println("main end")
	return
}
