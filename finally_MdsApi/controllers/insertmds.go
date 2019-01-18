package controllers

import (
	"MdsApi/libs"
	"MdsApi/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type InsertMdsController struct {
	beego.Controller
}


func (i *InsertMdsController) Post(){
	result := make(map[string]interface{})
	jsontomap := make(map[string]interface{})

	gameId := i.GetString("gameId")
	systemId := i.GetString("systemId")
	serverListStr := i.GetString("serverListStr")
	libs.NewLog().Debug("传入的数据是gameId:%s,systemId:%s,serverListStr:%s",gameId,systemId,serverListStr)

	var insertmds []interface{}
	var server []byte = []byte(serverListStr)
	err := json.Unmarshal(server, &jsontomap)
	if err != nil{
		libs.NewLog().Error("json转换map失败",err)
	}else{
		configId := jsontomap["configId"]
		gameAlias := jsontomap["gameAlias"]
		gameId := jsontomap["gameId"]
		platformId := jsontomap["platformId"]
		platformName := jsontomap["platformName"]
		platformAlias := jsontomap["platformAlias"]
		serverId := jsontomap["serverId"]
		serverName := jsontomap["serverName"]
		serverIp := jsontomap["serverIp"]
		unicomIp := jsontomap["unicomIp"]
		isCombined := jsontomap["isCombined"]
		isDeleted := jsontomap["isDeleted"]
		iszone := jsontomap["iszone"]
		serverOrder := jsontomap["serverOrder"]
		sumzone := jsontomap["sumzone"]
		insertmds = []interface{}{configId,gameAlias,gameId,platformId,platformName,platformAlias,serverId,serverName,serverIp,unicomIp,isCombined,isDeleted,iszone,serverOrder,sumzone}
	}
	sql := "insert into mds_server (configId,gameAlias,gameId,platformId,platformName,platformAlias,serverId,serverName,serverIp,unicomIp,isCombined,isDeleted,iszone,serverOrder,sumzone)\n" +
		" values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	dbname := beego.AppConfig.String(gameId+"mds")
	models.Init(dbname)
	o := orm.NewOrm()
	res,err := o.Raw(sql,insertmds).Exec()
	if err == nil{
		num,_ := res.RowsAffected()
		result["code"] = num
		result["msg"] = "插入数据成功！！！"
	}else{
		libs.NewLog().Error("faieds",err)
		result["msg"] = err
		result["code"] = "0"
	}

	i.Data["json"] = result
	i.ServeJSON()
	return


}