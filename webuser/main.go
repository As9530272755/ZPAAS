package main

import (
	"exec/config"
	"exec/route"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := ":8888"
	route.Register()

	driverName := "mysql"
	dsn := "root:root@tcp(10.0.0.10:3306)/hellodb?charset=utf8mb4&loc=Local&parseTime=true"
	config.ConnectDB(driverName, dsn)
	defer config.CloseDB()

	http.ListenAndServe(addr, nil)
}
