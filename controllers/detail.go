package controllers

import (
	"myblog/models"
	"strconv"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DetailController struct {
	beego.Controller
}

func (this *DetailController) Get() {

	typearticle := this.Ctx.Input.Param(":id")
	logs.Debug(typearticle)

	o := orm.NewOrm()
	article := new(models.Article)
	qs := o.QueryTable(article) // 返回 QuerySeter
	var articleitem models.Article

	id, err := strconv.Atoi(typearticle)
	logs.Debug(id)
	if err != nil {
		id = 1
	}
	if id < 1 {
		id = 1
	}
	qs.Filter("id", id).One(&articleitem)

	this.TplName = "detail.html"
	logs.Debug(articleitem)
	this.Data["article"] = articleitem
}
