package controllers

import (
	"github.com/astaxie/beego"
)

type HotKeyController struct {
	beego.Controller
}

func (h *HotKeyController) Get() {
	h.TplName = "demo/scan-hot-key.html"
}

