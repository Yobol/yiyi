package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "yiyi.me"
	c.Data["Email"] = "yobol1024@gmail.com"
	c.TplName = "index.tpl"
}
