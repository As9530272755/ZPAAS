package router

import (
	"ormweb/controller"

	"github.com/astaxie/beego"
)

func Register() {
	beego.Router("/", &controller.Getuser{}, "get:List")
	beego.Router("/user/add/", &controller.Adduser{}, "get:Add;post:Add")
	beego.Router("/user/delete", &controller.DeleteUser{}, "get:Delete")
	beego.Router("/user/edit", &controller.EditUser{}, "get:Edit;post:Edit")
	beego.Router("/login", &controller.Login{}, "get:Login;post:Login")
}
