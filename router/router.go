package router

import (
	"beegouser/controller"

	"github.com/astaxie/beego"
)

func Register() {
	beego.AutoRouter(&controller.AuthController{})
	beego.AutoRouter(&controller.UserController{})
}
