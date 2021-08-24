package controller

import (
	"fmt"

	"github.com/astaxie/beego"

	// "weborm/base"
	"weborm/models"
	"weborm/service"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	var errMsg string
	if user := c.GetSession("user"); user != nil {
		c.Redirect("/user/listuser", 302)
		return
	}

	if c.Ctx.Input.IsPost() {
		form := models.LoginForm{}
		if err := c.ParseForm(&form); err == nil {
			if user := service.Auth(&form); user != nil {
				fmt.Println("登陆成功")
				c.SetSession("user", user.Id)
				c.Redirect("/user/listuser", 302)
				return
			} else {
				errMsg = "提示：用户名或者密码错误"
			}
		}
	}

	c.Data["form"] = nil
	c.Data["error"] = errMsg
	c.TplName = "auth/login.html"
}
