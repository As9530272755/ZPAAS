package service

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"weborm/models"
)

// 查询用户
func GetUser() []*models.Webuser {
	ormer := orm.NewOrm()
	users := []*models.Webuser{}
	queryset := ormer.QueryTable(&models.Webuser{})
	queryset.All(&users)
	return users
}

// 添加用户
func AddUser(user *models.Webuser) error {
	Newpassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(Newpassword)
	fmt.Println("AddUser", user)
	ormer := orm.NewOrm()
	_, err := ormer.Insert(user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// 删除用户
func Delete(id int64) {
	ormer := orm.NewOrm()
	user := &models.Webuser{Id: id}
	ormer.Delete(user)
}

// 修改用户
func Edit(id int64, name, age, password, addr string, sex bool) {
	fmt.Println("edit service id = ", id)
	ormer := orm.NewOrm()
	user := &models.Webuser{Id: id}
	ormer.Read(user)
	user.Name = name
	user.Age = age
	user.Password = password
	user.Addr = addr
	user.Sex = sex
	ormer.Update(user)
}

// 在修改用户之前需要先查询该表
func Find(id int64) models.Webuser {
	ormer := orm.NewOrm()
	user := &models.Webuser{Id: id}
	ormer.Read(user)
	return *user
}

// 指定查询登录的 name 字段是否正确
func GetUserByName(name string) *models.Webuser {
	ormer := orm.NewOrm()
	user := &models.Webuser{Name: name}
	if err := ormer.Read(user, "name"); err == nil {
		return user
	}
	return nil
}

// 对用户输入的数据进行解析
func Auth(form *models.LoginForm) *models.Webuser {
	if user := GetUserByName(form.Username); user == nil {
		return nil
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err == nil {
			return user
		} else {
			return nil
		}
	}
}
