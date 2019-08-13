package main

import (
	_ "myblog/routers"
	"os"

	_ "myblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	MYSQL := os.Getenv("MYSQL")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", MYSQL)
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
