package base

import (
	"log"
	"ormweb/models"
	"ormweb/service"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type RequiredAuthController struct {
	beego.Controller
}

func (c *RequiredAuthController) Prepare() {
	user := c.GetSession("user")
	if user == nil {
		c.Redirect("/login", 302)
	}
}

func InitAdmin() {
	if user := service.GetUserByName("root"); user == nil {
		log.Println("创建管理员")
		newpassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), 5)
		adminUser := models.User{
			Name:     "root",
			Password: string(newpassword),
		}
		service.Add(&adminUser)
	}
}
