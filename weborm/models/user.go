package models

import "time"

type Webuser struct {
	Id       int64  `form:"id" orm:"column(id);auto"`
	Name     string `form:"name" orm:"size(32);"`
	Age      string `form:"age" orm:"size(32);"`
	Tel      string `form:"tel" orm:"size(32);"`
	Addr     string `form:"addr" orm:"type(text);"`
	Password string `form:"password" orm:"size(1024)"`
	Sex      bool
	Created  *time.Time `orm:"auto_now_add"`
	Updated  *time.Time `orm:"auto_now"`
	Deleted  *time.Time `orm:"null"`
}

func (u *Webuser) TableName() string {
	return "webuser"
}

// func Newwebuser(id int64, name, age, tel, addr, Password string, sex bool) *Webuser {
// 	return &Webuser{
// 		Id:       id,
// 		Name:     name,
// 		Age:      age,
// 		Tel:      tel,
// 		Addr:     addr,
// 		Password: Password,
// 		Sex:      sex,
// 	}
// }
