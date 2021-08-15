package controller

import (
	base "beegouser/base/controllers"
	"fmt"
)

type HomeController struct {
	base.RequiredAuthController
}

func (c *HomeController) Index() {
	fmt.Println("home")
	c.Ctx.WriteString("home")
}
