package controllers

import (
	"MdsApi/libs"
	"MdsApi/models"
	_ "encoding/json"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/dgrijalva/jwt-go"
	"strings"
	_ "time"
)

type OperaTionQuery struct {
	BaseController
}



func (q *OperaTionQuery) Post() {
	result := make(map[string]interface{})
	platDic := make(map[string]string)
	newplatDic := make(map[string]string)
	fields := q.GetString("fields")
	//切分fields字符串，判断查询字段的个数
	cut_fields := strings.Split(fields,",")
	count := len(cut_fields)
	platformAlias := q.GetString("platformAlias")
	serverId := q.GetString("serverId")
	platDic["serverId"] = serverId
	serverIp := q.GetString("serverIp")
	configId := q.GetString("configId")
	platDic["configId"] = configId
	combinedTo := q.GetString("combinedTo")
	platDic["combinedTo"] = combinedTo
	gameAlias := q.GetString("gameAlias")
	isCombined := q.GetString("isCombined")
	platDic["isCombined"] = isCombined

	for key,value := range platDic{
		if strings.TrimSpace(value) != ""{
			newplatDic[key] = value
		}
	}
	var logstr string
	logstr = "传入后端的数据有"
	for key,value := range newplatDic{
		logstr += ","+ key + ":" + value
	}
	if serverIp != ""{
		logstr += ",fields:%s,platformAlias:%s,gameAlias:%s,serverIp:%s"
		libs.NewLog().Debug(logstr,fields,platformAlias,gameAlias,serverIp)
	}else{
		logstr += ",fields:%s,platformAlias:%s,gameAlias:%s"
		libs.NewLog().Debug(logstr,fields,platformAlias,gameAlias)
	}

	var selectIds []string
	var sql string
	sql2 := "select " + fields
	if serverIp != ""{
		sql = sql2 + " from mds_server where platformAlias = ? and serverIp = ?"
		selectIds = []string{platformAlias,serverIp}
	}else{
		sql = sql2 + " from mds_server where platformAlias = ?"
		selectIds = []string{platformAlias}
	}
	for key,value := range newplatDic{
		sql += " and "+ key + " =" + value
	}

	dbname := beego.AppConfig.String(gameAlias+"mds")
	models.Init(dbname)
	o := orm.NewOrm()
	if count == 1{
		var maps []orm.Params
		res,err := o.Raw(sql,selectIds).Values(&maps)
		if err == nil && res > 0{
			result["code"] = "1"
			result["server_info"] = maps[0][fields]

		}else{
			result["code"] = "0"
			result["msg"] = "数据库执行错误,请正确输入参数"
			libs.NewLog().Error("数据库执行错误！！！,请正确输入参数",err)
			}
	}else{
		var list []orm.ParamsList
		res,err := o.Raw(sql,selectIds).ValuesList(&list)
		if err == nil && res > 0{
			result["code"] = "1"
			result["server_info"] = list[0]
		}else{
			result["code"] = "0"
			result["msg"] = err
			libs.NewLog().Error("数据库执行错误！！！,请正确输入参数",err)
		}
	}


	q.Data["json"] = result
	q.ServeJSON()
	return
}




