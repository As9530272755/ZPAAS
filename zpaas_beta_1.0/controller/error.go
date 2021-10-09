package controller

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.TplName = "error/err404.html"
}

func (c *ErrorController) Error500() {
	c.TplName = "error/err500.html"
}
