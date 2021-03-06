package models

import "time"

type User struct {
	Id       int64      `form:"id" orm:"column(id);auto" json:"id"`
	Name     string     `form:"name" orm:"size(32);"`
	Age      string     `form:"age" orm:"size(32);"`
	Tel      string     `form:"tel" orm:"size(32);"`
	Addr     string     `form:"addr" orm:"type(text);"`
	Password string     `form:"password" orm:"size(1024)" json:"-"`
	Sex      bool       `form:"sex"`
	Created  *time.Time `orm:"auto_now_add"`
	Updated  *time.Time `orm:"auto_now"`
	Deleted  *time.Time `orm:"null"`
}

func (u *User) TableName() string {
	return "webuser"
}
