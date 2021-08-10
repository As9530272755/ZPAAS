package main

import (
	"beegouser/config"
	"beegouser/router"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	driverName := "mysql"
	dsn := "root:root@tcp(10.0.0.10:3306)/hellodb?charset=utf8mb4&loc=Local&parseTime=true"
	addr := ":8000"

	config.LinkDb(driverName, dsn)
	defer config.CloseDb()
	router.Register()
	beego.Run(addr)
}
