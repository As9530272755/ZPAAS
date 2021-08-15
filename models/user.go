package models

type User struct {
	ID       int64  `form:"id"`
	Name     string `form:"name"`
	Age      string `form:"age"`
	Sex      bool   `form:"sex"`
	Addr     string `form:"addr"`
	Password string `form:"password"`
}

func NewUser(id int64, name, age string, sex bool, addr string) *User {
	return &User{ID: id, Name: name, Age: age, Sex: sex, Addr: addr}
}
