package controller

import "zpaas_beta_1.0/base"

type RootPath struct {
	base.RequiredAuthController
}

func (c *RootPath) Root() {
	c.TplName = "root/root.html"
}
