package router

import (
	"weborm/controller"

	"github.com/astaxie/beego"
)

func Register() {
	beego.AutoRouter(&controller.UserController{})
	beego.AutoRouter(&controller.AuthController{})
}
