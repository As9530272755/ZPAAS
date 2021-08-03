package models

type User struct {
	ID   int64
	Name string
	Age  string
	Sex  bool
	Addr string
}

func NewUser(id int64, name, age, addr string, sex bool) *User {
	return &User{
		ID:   id,
		Name: name,
		Age:  age,
		Sex:  sex,
		Addr: addr,
	}
}
