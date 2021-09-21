package main

import (
	"ormweb/base"
	"ormweb/models"
	"ormweb/router"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := ":8080"
	dsn := beego.AppConfig.String("dsn")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(&models.User{})
	orm.RunSyncdb("default", true, true)
	base.InitAdmin()

	router.Register()
	beego.Run(addr)
}
