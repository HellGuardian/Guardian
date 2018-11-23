package controllers

import (
	"github.com/astaxie/beego"
	"Guardian/models"
	"github.com/golang/glog"
	"fmt"
	"time"
	"strings"
	"github.com/tealeg/xlsx"
	"strconv"
)

type NoAuthController struct {
	beego.Controller
}

type DateStyle string
const YYYYMMDDHHMMSS = "yyyyMMddHHmmss"

func (n *NoAuthController) Get() {
	n.TplName = "demo/noauth_space.html"

	counts, err := models.GetSpaceNum()
	if err != nil {
		glog.Error("Get space num error")
		return
	}
	n.Data["Total"] = counts

	result, num, err := models.GetNoAuthSpace()
	if err != nil {
		glog.Error("Get no auth space error")
		return
	}

	n.Data["Num"] = num
	n.Data["NoAuthSpace"] = result

	fTotal := float64(counts)
	fNum := float64(num)
	n.Data["Per"] = fmt.Sprintf("%.2f", fNum / fTotal * 100)

	//excelName := time.Now().Format("2006-01-02 15:04:05")
	date := fmt.Sprintf(FormatDate(time.Now(), YYYYMMDDHHMMSS))
	// save excel fle
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "集群ID"
	cell = row.AddCell()
	cell.Value = "集群名"
	cell = row.AddCell()
	cell.Value = "集群所有者"
	cell = row.AddCell()
	cell.Value = "一级部门"
	cell = row.AddCell()
	cell.Value = "二级部门"
	cell = row.AddCell()
	cell.Value = "erp"
	cell = row.AddCell()
	cell.Value = "授权邮箱"

	for _, v := range result{
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = strconv.Itoa(v.Id)
		cell = row.AddCell()
		cell.Value = v.Name
		cell = row.AddCell()
		cell.Value = v.Role
		cell = row.AddCell()
		cell.Value = v.Dep1
		cell = row.AddCell()
		cell.Value = v.Dep2
		cell = row.AddCell()
		cell.Value = v.Erp
		cell = row.AddCell()
		cell.Value = v.Mail
	}
	excelName := fmt.Sprintf("%s%s%s","NoAuth-", date, ".xlsx")
	err = file.Save(excelName)
	if err != nil {
		glog.Error("save excel file error")
		return
	}
	n.Data["excelName"] = excelName
}

func FormatDate(date time.Time, dateStyle DateStyle) string {
	layout := string(dateStyle)
	layout = strings.Replace(layout, "yyyy", "2006", 1)
	layout = strings.Replace(layout, "yy", "06", 1)
	layout = strings.Replace(layout, "MM", "01", 1)
	layout = strings.Replace(layout, "dd", "02", 1)
	layout = strings.Replace(layout, "HH", "15", 1)
	layout = strings.Replace(layout, "mm", "04", 1)
	layout = strings.Replace(layout, "ss", "05", 1)
	layout = strings.Replace(layout, "SSS", "000", -1)

	return date.Format(layout)
}