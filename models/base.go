package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
)

var engine *xorm.Engine

func InitDB() (err error) {
	dbUrl := beego.AppConfig.String("dbUrl")
	if dbUrl == "" {
		glog.Error("mysql url is null.")
		return err
	}
	engine, err = xorm.NewEngine("mysql", dbUrl)
	if err != nil {
		glog.Error("new mysql engine error")
		return err
	}

	maxIdleConn, err := beego.AppConfig.Int("maxIdleConns")
	if err != nil {
		glog.Error("get maxIdleConns config error")
		return err
	}
	engine.SetMaxIdleConns(maxIdleConn)
	maxOpenConn, err := beego.AppConfig.Int("maxOpenConns")
	if err != nil {
		glog.Error("get maxOpenConns config error")
		return err
	}
	engine.SetMaxOpenConns(maxOpenConn)

	if err := engine.Ping(); err != nil {
		glog.Error("MySql ping error.")
		return err
	}
	return err
}
