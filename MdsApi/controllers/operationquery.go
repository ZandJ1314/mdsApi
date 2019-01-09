package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type OperaTionQuery struct {
	beego.Controller
}

func (q *OperaTionQuery) Get() map[string]string {
	result := make(map[string]string)
	//fmt.Println("testters")
	fields := q.GetString("fields")
	//切分fields字符串，判断查询字段的个数
	cut_fields := strings.Split(fields,",")
	count := len(cut_fields)
	platformAlias := q.GetString("platformAlias")
	serverId := q.GetString("serverId")
	serverIp := q.GetString("serverIp")
	configId := q.GetString("configId")
	combinedTo := q.GetString("combinedTo")
	gameAlias := q.GetString("gameAlias")
	isCombined := q.GetString("isCombined")
	time := q.GetString("time")
	sign := q.GetString("sign")
	fmt.Println(time,sign)
	//gameAlias := "lwtz"
	//platformAlias := "37_lwtz"
	//serverId := "50004"
	//serverIp := "47.97.23.94"
	//configId := "57"
	//combinedTo := "50001"
	//isCombined := "1"
	//查询龙纹天尊数据库

	if gameAlias == "lwtz"{
		o := orm.NewOrm()
		if count == 1{
			var maps []orm.Params
			selectIds := []string{fields,platformAlias,serverId,serverIp,configId,combinedTo,isCombined}
			sql := "select ? from mds_server where platformAlias = ? and serverId = ? and serverIp=? and configId=? and combinedTo=? and isCombined=?"
			res,err := o.Raw(sql,selectIds).Values(&maps)
			if err == nil && res > 0{
				result["code"] = "1"
				result["server_info"] = ""

			}else{
				result["code"] = "0"
				result["msg"] = "数据库执行错误"
			}
		}else{
			var list []orm.ParamsList
			selectIds := []string{fields,platformAlias,serverId,serverIp,configId,combinedTo,isCombined}
			sql := "select ? from mds_server where platformAlias = ? and serverId = ? and serverIp=? and configId=? and combinedTo=? and isCombined=?"
			res,err := o.Raw(sql,selectIds).ValuesList(&list)
			if err == nil && res > 0{
				result["code"] = "1"
				result["server_info"] = ""
			}else{
				result["code"] = "0"
				result["msg"] = "数据库执行错误"
			}
		}



	}else{
		o1 := orm.NewOrm()
		o1.Using("db2")
		var maps []orm.Params
		selectIds := []string{fields,platformAlias,serverId,serverIp,configId,combinedTo,isCombined}
		sql := "select ? from mds_server where platformAlias = ? and serverId = ? and serverIp=? and configId=? and combinedTo=? and isCombined=?"
		res,err := o1.Raw(sql,selectIds).Values(&maps)
		fmt.Println(res,err)
		if err == nil{
			fmt.Println(maps[0][fields])
		}
	}


	return result
}
