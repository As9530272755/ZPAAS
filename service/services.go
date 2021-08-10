package service

import (
	"beegouser/config"
	"beegouser/models"
	"fmt"
	"log"
)

func GetUser() []*models.User {
	users := make([]*models.User, 0)
	rows, err := config.DB.Query("select id,name,age,sex,addr from WEBuser;")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		var (
			ID   int64
			Name string
			Age  string
			Sex  bool
			Addr string
		)

		err := rows.Scan(&ID, &Name, &Age, &Sex, &Addr)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		users = append(users, models.NewUser(ID, Name, Age, Sex, Addr))
	}
	return users
}

func AddUser(name, age, addr string, sex bool) {
	fmt.Println("AddUser", name, age, addr, sex)
	_, err := config.DB.Exec("insert into WEBuser(name , age , sex ,addr,created_at,updated_at) values(? , ? , ? , ?,now(),now())", name, age, sex, addr)
	if err != nil {
		log.Println(err)
		return
	}
}

func DeleteUser(id int64) {
	_, err := config.DB.Exec("delete from WEBuser where id=?;", id)
	if err != nil {
		log.Println(err)
		return
	}
}

func Edit(id int64, name, age, addr string, sex bool) {
	fmt.Println("EditUser", id, name, age, addr, sex)
	_, err := config.DB.Exec("update WEBuser set name=?,age=?,addr=?,sex=? where id=?;", name, age, addr, sex, id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(id)
}

func IDFindUser(Id int64) (models.User, error) {
	// users := models.User{}

	row := config.DB.QueryRow("select name,age,sex,addr from WEBuser where id=?;", Id)
	var (
		name string
		age  string
		sex  bool
		addr string
	)

	err := row.Scan(&name, &age, &sex, &addr)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	} else {
		users := models.User{
			ID:   Id,
			Name: name,
			Age:  age,
			Sex:  sex,
			Addr: addr,
		}
		return users, nil
	}
}
