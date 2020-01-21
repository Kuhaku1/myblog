package controllers

import (
	"fmt"
	"myblog/models"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ShowController struct {
	beego.Controller
}

func (this *ShowController) Get() {

	typearticle := this.GetString("type", "other")
	logs.Debug(typearticle)
	o := orm.NewOrm()
	article := new(models.Article)
	qs := o.QueryTable(article) // 返回 QuerySeter
	var articles []*models.Article
	qs.Filter("Category", typearticle).All(&articles)
	for _, v := range articles {
		fmt.Println(*v)
	}
	this.TplName = "showlist.html"
	this.Data["articles"] = articles
}
