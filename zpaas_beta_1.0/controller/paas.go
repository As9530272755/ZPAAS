package controller

import "zpaas_beta_1.0/base"

type Paas struct {
	base.RequiredAuthController
}

func (c *Paas) Paas() {
	c.Data["navKey"] = "paas"
	c.TplName = "paas/paas.html"
}
