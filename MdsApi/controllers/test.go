package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TestQuery struct {
	beego.Controller
}


func (t *TestQuery) Get() map[string]string {
	result := make(map[string]string)
	//gameAlias := "lwtz"
	platformAlias := "37_lwtz"
	serverId := "50004"
	serverIp := "47.97.23.94"
	configId := "57"
	combinedTo := "50001"
	isCombined := "1"
	o := orm.NewOrm()
	var maps []orm.Params
	selectIds := []string{platformAlias,serverId,serverIp,configId,combinedTo,isCombined}
	sql := "select javaDir from mds_server where platformAlias = ? and serverId = ? and serverIp=? and configId=? and combinedTo=? and isCombined=?"
	res,err := o.Raw(sql,selectIds).Values(&maps)
	if err == nil && res >0 {
		fmt.Println(maps[0]["javaDir"])
		result["code"] = "0"
	}
	return result
}
