package controller

import (
	"fmt"
	"ormweb/models"
	"ormweb/service"

	"github.com/astaxie/beego"
)

type Login struct {
	beego.Controller
}

func (c *Login) Login() {
	var errMsg string

	if user := c.GetSession("user"); user != nil {
		c.Redirect("/", 302)
		return
	}

	if c.Ctx.Input.IsPost() {
		authuser := models.User{}
		c.ParseForm(&authuser)
		fmt.Println(authuser)
		if user := service.Auth(&authuser); user != nil {
			fmt.Println("登陆成功")
			c.SetSession("user", user.Id)
			c.Redirect("/", 302)
			return
		} else {
			errMsg = "提示：用户名或密码输入错误"
		}
	}

	c.Data["user"] = nil
	c.Data["error"] = errMsg
	c.TplName = "login/login.html"
}
