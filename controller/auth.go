package controller

import (
	"beegouser/errors"
	"beegouser/models"
	"beegouser/service"
	"fmt"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

// 打开页面和点击登录都交给 Login 来处理
func (c *AuthController) Login() {
	// 登录成功就不让在访问 login 的页面
	if user := c.GetSession("user"); user != nil {
		c.Redirect("/user/listuser", 302)
		return
	}

	// 有的时候错误信息可能是多个,所以这里定义一个 errors 的map，并且 value 是[]
	errors := errors.NewErrors()

	form := models.LoginForm{}
	// 如果是 post 请求就登录
	if c.Ctx.Input.IsPost() {
		// 将用户提交的数据。解析我们的 loginform 结构体中
		if err := c.ParseForm(&form); err == nil {
			if user := service.Auth(&form); user != nil {
				// 登录成功就跳转值 /user/listuser 界面
				fmt.Println("登陆成功")
				// 登录成功存储 session 读取 user.ID
				c.SetSession("user", user.ID)
				c.Redirect("/user/listuser", 302)
				return
			} else {
				errors.AddError("default", "提示：用户名或密码输入错误")
			}
		}
	}

	// 用户输入数据回显
	c.Data["form"] = form
	c.Data["errors"] = errors
	// 如果不是 post 打开登陆页面
	c.TplName = "auth/login.html"
}

func (c *AuthController) Logout() {
	c.DestroySession()
	c.Redirect("/auth/login", 302)
}
