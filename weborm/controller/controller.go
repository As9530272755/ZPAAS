package controller

import (
	"fmt"
	"log"
	"weborm/base"
	"weborm/models"
	"weborm/service"
)

// 定义user 控制器
type UserController struct {
	base.RequiredAuthController
}

// 处理显示所有 user 信息
func (c *UserController) ListUser() {
	c.Data["user"] = service.GetUser()
	c.TplName = "user/listuser.html"
}

// 添加 user
func (c *UserController) Add() {
	if c.Ctx.Input.IsPost() {
		var user models.Webuser
		fmt.Println("User", user)
		c.ParseForm(&user)
		fmt.Println("UserController", user)
		service.AddUser(&user)
		c.Redirect("/user/listuser.html", 302)
	} else {
		c.TplName = "user/add.html"
	}
}

// 删除用户
func (c *UserController) Delete() {
	if id, err := c.GetInt64("id"); err == nil {
		service.Delete(id)
		c.Redirect("/user/listuser.html", 302)
	} else {
		log.Println(err)
		return
	}
}

// 编辑用户
func (c *UserController) Edit() {
	type users struct {
		Name     string `form:"name"`
		Age      string `form:"age"`
		Password string `form:"password"`
		Sex      bool   `form:"sex"`
		Addr     string `form:"addr"`
	}
	id, _ := c.GetInt64("id")
	if c.Ctx.Input.IsPost() {
		fmt.Println("post", id)
		var user users
		// 解析用户在 wbe 页面传递的数据
		c.ParseForm(&user)
		// 传递编辑参数
		service.Edit(id, user.Name, user.Age, user.Password, user.Addr, user.Sex)
		c.Redirect("/user/listuser.html", 302)
	} else {
		fmt.Println(id)
		// 调用 find 查找对应 id 的信息
		c.Data["user"] = service.Find(id)
		c.TplName = "user/edit.html"
	}
}
