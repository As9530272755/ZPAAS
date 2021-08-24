package main

import (
	"log"
	"weborm/models"
	"weborm/router"
	"weborm/service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	addr := ":8888"
	// 通过 AppConfig 读取 conf 目录下的 app.conf 配置文件中的 dsn
	dsn := beego.AppConfig.String("dsn")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(&models.Webuser{})
	orm.RunSyncdb("default", false, true)

	// 初始化管理员用户
	if user := service.GetUserByName("admin"); user == nil {
		log.Println("创建管理员用户")
		Newpassword, _ := bcrypt.GenerateFromPassword([]byte("123@456"), 10)
		adminuser := models.Webuser{Name: "admin", Password: string(Newpassword)}
		service.AddUser(&adminuser)
	}

	router.Register()
	beego.Run(addr)
}
