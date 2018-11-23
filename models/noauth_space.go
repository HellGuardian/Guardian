package models

import (
	"github.com/golang/glog"
	"strconv"
)

type NoAuthSpaceInfo struct {
	Id		int
	Name	string
	Role	string
	Dep1	string
	Dep2	string
	Erp		string
	Mail	string
}

type Space struct {
	Id		int
}

func GetNoAuthSpace() (noAuth []*NoAuthSpaceInfo, num int, err error) {
	querySql := "select  a.name ,a.id,b.role,b.department1,b.department2,a.erp,a.mail  from (select a.name as name,a.space_id as ID,erp,mail from (select space.name,space_user.space_id,group_concat(space_user.erp) as erp from  space,space_user where space.id=space_user.space_id  and space.id not in  (select space_id from space_extend where config_commands like  '%config set requirepass%' and   config_commands not  like  '%config set requirepass \"\"%') group by space_id ) a  join (select space_user.space_id,user.department1,user.department2,group_concat(user.mail)  as mail from  user,space_user where space_user.erp=user.erp group by space_id )  b on a.space_id= b.space_id ) a  join (select space_id ,user.erp as role, department1 as department1, ifnull(department1,department2) as department2 from space_user, user  where  user.erp=space_user.erp  and role=0)  b on a.id= b.space_id  order by b.role"

	result, err := engine.Query(querySql)
	if err != nil {
		glog.Error("Query no auth space error")
		return
	} else if len(result) == 0 {
		glog.Error("Query admin result is emtpy.")
		return
	}

	num = len(result)

	for _, v := range result {
		var Info *NoAuthSpaceInfo
		Info = new(NoAuthSpaceInfo)
		Info.Name = string(v["name"])
		Info.Role = string(v["role"])
		Info.Mail = string(v["mail"])
		Info.Erp = string(v["erp"])
		Info.Dep1 = string(v["department1"])
		Info.Dep2 = string(v["department2"])
		Info.Id, _ = strconv.Atoi(string(v["id"]))
		noAuth = append(noAuth, Info)
	}

	return noAuth, num, err
}

func GetSpaceNum() (num int64, err error) {
	space := new(Space)
	num, err = engine.Count(space)
	if err != nil {
		glog.Error("Query space num error")
		return num, err
	}
	return num, err
}