package controllers

import (
	"encoding/json"
	"fmt"
	"myblog/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type artcsjson struct {
	Topic string `json:"topic"`
	Types string `json:"type"`
	Value string `json:"value"`
}

type CreateController struct {
	beego.Controller
}

func (c *CreateController) Get() {

	c.TplName = "create.html"

}
func (c *CreateController) Post() {

	var ob artcsjson
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	if err != nil {
		fmt.Println(err.Error())
	}

	o := orm.NewOrm()
	var article models.Article
	article.Category = ob.Types
	article.Title = ob.Topic
	article.Content = ob.Value
	article.Created = time.Now()
	_, err = o.Insert(&article)
	if err != nil {
		c.Ctx.WriteString("error")
	} else {
		c.Ctx.WriteString("ok")
	}

}
