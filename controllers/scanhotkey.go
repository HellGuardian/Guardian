package controllers

import (
	"github.com/astaxie/beego"
	"Guardian/models"
	"github.com/golang/glog"
	"strings"
)

type HotKeyController struct {
	beego.Controller
}

func (h *HotKeyController) Get() {
	h.TplName = "demo/scan-hot-key.html"

	result, err := models.GetAdminInfo()
	if err != nil {
		glog.Error("Get admin table error.")
	}
	h.Data["AdminInfo"] = result

	ip := strings.TrimSpace(h.GetString("IpPort"))
	pwd := strings.TrimSpace(h.GetString("PassWord"))

	h.Data["Ip"] = ip
	h.Data["Pwd"] = pwd
}
