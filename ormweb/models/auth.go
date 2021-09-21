package models

type Auth struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}
