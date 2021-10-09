package controller

import "zpaas_beta_1.0/base"

type Prom struct {
	base.RequiredAuthController
}

func (c *Prom) Prom() {
	c.Data["navKey"] = "prometheus"
	c.TplName = "prom/prom.html"
}
