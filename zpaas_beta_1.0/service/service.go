package service

import (
	"fmt"

	"zpaas_beta_1.0/models"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

func Get() []*models.User {
	users := []*models.User{}
	ormer := orm.NewOrm()
	querset := ormer.QueryTable(&models.User{})
	querset.All(&users)
	return users
}

func Add(user *models.User) {
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 3)
	user.Password = string(newPassword)
	ormer := orm.NewOrm()
	fmt.Println(user)
	fmt.Println(ormer.Insert(user))
}

func Delete(id int64) {
	ormer := orm.NewOrm()
	user := &models.User{
		Id: id,
	}
	ormer.Delete(user)
}

func Find(id int64) *models.User {
	ormer := orm.NewOrm()
	ormer.QueryTable(&models.User{})
	user := &models.User{Id: id}
	ormer.Read(user)
	return user
}

func Edit(id int64, name, age, password, addr, tel string, sex bool) {
	ormer := orm.NewOrm()
	user := &models.User{Id: id}
	ormer.Read(user)
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 3)
	user.Name = name
	user.Age = age
	user.Password = string(newPassword)
	user.Addr = addr
	user.Tel = tel
	user.Sex = sex
	fmt.Println(user)
	ormer.Update(user)
}

func GetUserByName(name string) *models.User {

	ormer := orm.NewOrm()
	user := &models.User{Name: name}
	if err := ormer.Read(user, "name"); err == nil {
		return user
	}
	return nil
}

// 通过 orm 查询当前 id 的user 信息
func GetById(pk int64) *models.User {
	ormer := orm.NewOrm()
	ormer.QueryTable(&models.User{})
	user := &models.User{
		Id: pk,
	}
	if err := ormer.Read(user); err == nil {
		return user
	}
	return nil
}

func Auth(Auth *models.User) *models.User {
	if user := GetUserByName(Auth.Name); user == nil {
		return nil
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Auth.Password)); err == nil {
			return user
		} else {
			return nil
		}

	}
}
