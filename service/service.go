package service

import (
	"exec/config"
	"exec/models"
	"fmt"
	"log"
)

func ListUser() []*models.User {
	users := make([]*models.User, 0)
	rows, err := config.DB.Query("select id,name,age,sex,addr from WEBuser; ")
	if err != nil {
		log.Println(err)
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

		if err := rows.Scan(&ID, &Name, &Age, &Sex, &Addr); err != nil {
			log.Println(err)
			return nil
		}
		users = append(users, models.NewUser(ID, Name, Age, Addr, Sex))
	}
	return users
}

func AddUser(name, age, addr string, sex bool) {
	_, err := config.DB.Exec("insert into WEBuser(name , age , sex ,addr,created_at,updated_at) values(? , ? , ? , ?,now(),now())", name, age, sex, addr)
	if err != nil {
		log.Println(err)
		return
	}
}

func DeleteUser(id int64) {
	_, err := config.DB.Exec("delete from WEBuser where id=?", id)
	if err != nil {
		log.Println(err)
		return
	}
}

func IDFindUser(Id int64) (models.User, error) {
	var user = models.User{}
	sql := `
    SELECT id , name , age, sex ,addr  FROM WEBuser WHERE id = ?
    `
	var (
		id   int64
		name string
		age  string
		sex  bool
		addr string
	)

	err := config.DB.QueryRow(sql, Id).Scan(&id, &name, &age, &sex, &addr)
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	} else {
		user = models.User{
			ID:   id,
			Name: name,
			Age:  age,
			Sex:  sex,
			Addr: addr,
		}
		// fmt.Println(ID, Name, Age, Sex, Addr)
	}
	return user, nil
}

func Edit(id int64, name, age, addr string, sex bool) error {
	var err error
	sql := `update WEBuser set name= ? ,age= ?,addr= ? ,sex= ? where id= ?;`

	_, err = config.DB.Exec(sql, name, age, addr, sex, id)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("mod user %v LastInsertID and RowsAffected: ", id)
	return nil
}
