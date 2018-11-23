package models

import (
	"github.com/golang/glog"
	"strconv"
)

type AdminInfo struct {
	Erp		string
	Mail	string
	Tel		string
	Role	int
}

func GetAdminInfo() (adminInfo []*AdminInfo, err error) {
	adminSql := "select * from admin"
	result, err := engine.Query(adminSql)
	if err != nil {
		glog.Error("Query admin error")
		return
	} else if len(result) == 0 {
		glog.Error("Query admin result is emtpy.")
		return
	}

	for _, v := range result {
		var Info *AdminInfo
		Info = new(AdminInfo)
		Info.Erp = string(v["erp"])
		Info.Mail = string(v["mail"])
		Info.Tel = string(v["tel"])
		Info.Role, _ = strconv.Atoi(string(v["ROLE"]))
		adminInfo = append(adminInfo, Info)
	}

	return adminInfo, err
}
