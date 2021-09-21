package controller

import (
	"ormweb/base"
	"ormweb/models"
	"ormweb/service"
)

type Getuser struct {
	base.RequiredAuthController
}

func (c *Getuser) List() {
	c.Data["user"] = service.Get()
	c.TplName = "user/list.html"
}

type Adduser struct {
	base.RequiredAuthController
}

func (c *Adduser) Add() {
	if c.Ctx.Input.IsPost() {
		var user models.User
		c.ParseForm(&user)
		service.Add(&user)
		c.Redirect("/", 302)
	} else {
		c.TplName = "user/add.html"
	}
}

type DeleteUser struct {
	base.RequiredAuthController
}

func (c *DeleteUser) Delete() {
	id, _ := c.GetInt64("id")
	service.Delete(id)
	c.Redirect("/", 302)
}

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
		c.Redirect("/", 302)
	} else {
		c.Data["user"] = service.Find(id)
		c.TplName = "user/edit.html"
	}
}
