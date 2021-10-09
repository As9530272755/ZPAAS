package main

import (
	"zpaas_beta_1.0/base"
	"zpaas_beta_1.0/router"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 开启日志取消注释
	// base.Log()

	addr := beego.AppConfig.String("web::PORT")

	// 注册数据库
	base.DbRegister()

	// 注册路由
	router.Register()

	// 是否开启日志，如开启日志取消注释即可
	// base.Log()

	beego.Run(addr)
}
