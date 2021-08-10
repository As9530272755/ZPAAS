package controller

import (
	"beegouser/models"
	"beegouser/service"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) ListUser() {
	c.Data["user"] = service.GetUser()
	c.TplName = "users/user.html"
}

func (c *UserController) Add() {
	if c.Ctx.Input.IsPost() {
		var form models.User
		c.ParseForm(&form)
		service.AddUser(form.Name, form.Age, form.Addr, form.Sex)
		c.Redirect("/user/listuser", 302)
	} else {
		c.TplName = "users/add.html"
	}
}

func (c *UserController) Delete() {
	if id, err := c.GetInt64("id"); err == nil {
		service.DeleteUser(id)
		c.Redirect("/user/listuser", 302)
	} else {
		log.Println(err)
		return
	}
}

func (c *UserController) Edit() {
	type users struct {
		Name string `form:"name"`
		Age  string `form:"age"`
		Sex  bool   `form:"sex"`
		Addr string `form:"addr"`
	}
	id, _ := c.GetInt64("id")
	if c.Ctx.Input.IsPost() {
		fmt.Println("post", id)
		var form users
		c.ParseForm(&form)
		fmt.Println("form", form)
		service.Edit(id, form.Name, form.Age, form.Addr, form.Sex)
		c.Redirect("/user/listuser", 302)
	} else {
		fmt.Println(id)
		user, _ := service.IDFindUser(id)
		//users = user
		c.Data["user"] = user
		c.TplName = "users/edit.html"
	}
}
