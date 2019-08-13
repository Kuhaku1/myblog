package controllers

import (
	"github.com/astaxie/beego"
)

type CreateController struct {
	beego.Controller
}

func (c *CreateController) Get() {

	c.TplName = "create.html"

}
