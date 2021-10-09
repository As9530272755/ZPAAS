package controller

import (
	"zpaas_beta_1.0/base"
	"zpaas_beta_1.0/models"
	"zpaas_beta_1.0/service"

	"github.com/astaxie/beego"
)

type Getuser struct {
	base.RequiredAuthController
}

func (c *Getuser) List() {
	// 读取 flash 消息
	beego.ReadFromRequest(&c.Controller)

	c.Data["navKey"] = "user"
	c.Data["user"] = service.Get()
	c.TplName = "user/list.html"
}

// 定义 json 结构体
type GetJson struct {
	base.RequiredAuthController
}

// 定义 url
func (c *GetJson) UserJson() {
	// c.Data["json"] 的数据通过 service.Get() 函数获取
	c.Data["json"] = service.Get()
	c.ServeJSON()
}

// 添加用户
type Adduser struct {
	base.RequiredAuthController
}

func (c *Adduser) Add() {
	c.Data["navKey"] = "adduser"
	if c.Ctx.Input.IsPost() {
		var user models.User
		c.ParseForm(&user)
		service.Add(&user)

		// flash 消息回闪
		flash := beego.NewFlash()
		flash.Set("success", "添加用户成功!")
		flash.Store(&c.Controller)

		c.Redirect("/user/", 302)
	} else {
		c.TplName = "user/add.html"
	}
}

// 删除用户
type DeleteUser struct {
	base.RequiredAuthController
}

func (c *DeleteUser) Delete() {
	id, _ := c.GetInt64("id")
	service.Delete(id)
	// 定义 flash
	flash := beego.NewFlash()
	flash.Set("delete", "删除用户成功")
	flash.Store(&c.Controller)

	c.Redirect("/user/", 302)
}

// 编辑用户
type EditUser struct {
	base.RequiredAuthController
}

func (c *EditUser) Edit() {
	type Newuser struct {
		Name     string `form:"name"`
		Age      string `form:"age"`
		Password string `form:"password"`
		Sex      bool   `form:"sex"`
		Addr     string `form:"addr"`
		Tel      string `form:"tel"`
	}
	id, _ := c.GetInt64("id")
	if c.Ctx.Input.IsPost() {
		var user Newuser
		c.ParseForm(&user)
		service.Edit(id, user.Name, user.Age, user.Password, user.Addr, user.Tel, user.Sex)

		// 添加 flash 消息
		flash := beego.NewFlash()
		flash.Set("success", "编辑用户成功")
		flash.Store(&c.Controller)
		c.Redirect("/user", 302)
	} else {
		c.Data["user"] = service.Find(id)
		c.TplName = "user/edit.html"
	}
}
