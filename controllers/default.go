package controllers

import (
	"fmt"
	"myblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	// var article models.Article
	// article.Category = "动漫"
	// article.Title = "第一条数据"
	// article.Content = "第一条数据内容"
	// article.Created = time.Now()

	// id, err := o.Insert(&article)
	// if err == nil {
	// 	fmt.Println(id)
	// }

	article := new(models.Article)
	qs := o.QueryTable(article) // 返回 QuerySeter
	var articles []*models.Article
	qs.Limit(10).All(&articles)
	for _, v := range articles {
		fmt.Println(*v)
	}
	c.TplName = "index.html"
	c.Data["articles"] = articles

}
