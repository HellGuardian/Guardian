package controllers

import (
	"github.com/astaxie/beego"
)

type WelcomeController struct {
	beego.Controller
}

func (w *WelcomeController) Get() {
	w.TplName = "demo/welcome.html"
}
