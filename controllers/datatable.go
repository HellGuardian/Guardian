package controllers

import (
	"github.com/astaxie/beego"
)

type DataTableController struct {
	beego.Controller
}

func (d *DataTableController) Get() {
	d.TplName = "demo/data-table.html"
}
