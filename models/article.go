package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id       int    `orm:"pk"`
	Category string `orm:"null"`
	Title    string
	Content  string
	Created  time.Time
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Article))
}
