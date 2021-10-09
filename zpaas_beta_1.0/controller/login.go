package controller

import (
	"zpaas_beta_1.0/models"
	"zpaas_beta_1.0/service"

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
		if user := service.Auth(&authuser); user != nil {
			c.SetSession("user", user.Id)
			c.Redirect("/", 302)
			beego.Informational(authuser.Name, "登陆成功")
			return
		} else {
			errMsg = "提示：用户名或密码输入错误!"
		}
	}

	c.Data["user"] = nil
	c.Data["error"] = errMsg
	c.TplName = "login/login.html"
}

type Outlogin struct {
	beego.Controller
}

func (c *Outlogin) Out() {
	c.DestroySession()
	c.Redirect("/login", 302)
}
