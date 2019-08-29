package routers

import (
	"myblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// 创建
	beego.Router("/create", &controllers.CreateController{})
	beego.Router("/save", &controllers.CreateController{})
	// 查看
	beego.Router("/show", &controllers.ShowController{})
	beego.Router("/detail/:id", &controllers.DetailController{})
}
