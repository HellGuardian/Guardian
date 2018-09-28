package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "tanxiaowei@jd.com"
	c.Data["Email"] = "tanxiaowei@jd.com"
	c.TplName = "index.html"
}