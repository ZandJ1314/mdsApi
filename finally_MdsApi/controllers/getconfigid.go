package controllers

import (
	"MdsApi/libs"
	"MdsApi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type GetConfigId struct {
	beego.Controller
}

func (g *GetConfigId) Post(){
	result := make(map[string]interface{})
	platDic := make(map[string]string)
	newplatDic := make(map[string]string)
	platformAlias := g.GetString("platformAlias")
	serverId := g.GetString("serverId")
	platDic["serverId"] = serverId
	serverIp := g.GetString("serverIp")
	combinedTo := g.GetString("combinedTo")
	platDic["combinedTo"] = combinedTo
	gameAlias := g.GetString("gameAlias")
	isCombined := g.GetString("isCombined")
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
		logstr += ",platformAlias:%s,gameAlias:%s,serverIp:%s"
		libs.NewLog().Debug(logstr,platformAlias,gameAlias,serverIp)
	}else{
		logstr += ",platformAlias:%s,gameAlias:%s"
		libs.NewLog().Debug(logstr,platformAlias,gameAlias)
	}

	var selectIds []string
	var sql string
	if serverIp != ""{
		sql = "select configId from mds_server where platformAlias = ? and serverIp = ?"
		selectIds = []string{platformAlias,serverIp}
	}else{
		sql = "select configId from mds_server where platformAlias = ?"
		selectIds = []string{platformAlias}
	}
	for key,value := range newplatDic{
		sql += " and "+ key + " =" + value
	}

	dbname := beego.AppConfig.String(gameAlias+"mds")
	models.Init(dbname)
	o := orm.NewOrm()
	var maps []orm.Params
	res,err := o.Raw(sql,selectIds).Values(&maps)
	if err == nil && res > 0{
		result["code"] = "1"
		result["count"] = maps[0]["configId"]

	}else{
		result["code"] = "0"
		result["msg"] = "数据库执行错误,请正确输入参数"
		libs.NewLog().Error("数据库执行错误！！！,请正确输入参数",err)
	}

	g.Data["json"] = result
	g.ServeJSON()
	return


}


