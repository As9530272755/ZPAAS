package controller

import "zpaas_beta_1.0/base"

type WebShell struct {
	base.RequiredAuthController
}

func (c *WebShell) WebShell() {
	c.Data["navKey"] = "webshell"
	c.TplName = "web/webshell.html"
}
