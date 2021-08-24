package base

import "github.com/astaxie/beego"

type RequiredAuthController struct {
	beego.Controller
}

func (c *RequiredAuthController) Prepare() {
	user := c.GetSession("user")
	if user == nil {
		c.Redirect("/auth/login", 302)
	}
}
