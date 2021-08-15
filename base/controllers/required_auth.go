package controllers

import "github.com/astaxie/beego"

type RequiredAuthController struct {
	beego.Controller
}

func (c *RequiredAuthController) Prepare() {
	user := c.GetSession("user")
	if user == nil {
		// 如果 user 为 nil 就是未登录，就重定向到登录页面
		c.Redirect("/auth/login", 302)
	}
}
