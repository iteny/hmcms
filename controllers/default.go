package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me11111111111"
	c.Data["Email"] = "astaxie@gmail.com1111111111"
	c.TplName = "index.tpl"
}
