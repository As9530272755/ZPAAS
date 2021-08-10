package models

type User struct {
	ID   int64  `form:"id"`
	Name string `form:"name"`
	Age  string `form:"age"`
	Sex  bool   `form:"sex"`
	Addr string `form:"addr"`
}

func NewUser(id int64, name, age string, sex bool, addr string) *User {
	return &User{id, name, age, sex, addr}
}
