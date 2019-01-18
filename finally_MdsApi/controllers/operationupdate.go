package controllers

import (
	"MdsApi/libs"
	"MdsApi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type UpdateController struct {
	BaseController
}

func (u *UpdateController) Post(){
	result := make(map[string]interface{})
	platDic := make(map[string]string)
	newplatDic := make(map[string]string)
	fields := u.GetString("fields")
	values := u.GetString("values")
	//切分fields字符串，判断查询字段的个数
	cut_fields := strings.Split(fields,",")
	cut_values := strings.Split(values,",")
	count := len(cut_fields)
	platformAlias := u.GetString("platformAlias")
	serverId := u.GetString("serverId")
	platDic["serverId"] = serverId
	serverIp := u.GetString("serverIp")
	configId := u.GetString("configId")
	platDic["configId"] = configId
	combinedTo := u.GetString("combinedTo")
	platDic["combinedTo"] = combinedTo
	gameAlias := u.GetString("gameAlias")
	isCombined := u.GetString("isCombined")
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
	sql2 := "update mds_server set "

	if count == 1{
		sql2 += fields + "=" + values
	}else{
		for i:=0; i<count;i++{
			sql2 += cut_fields[i] + "=" + cut_values[i]
		}
	}
	if serverIp != ""{
		sql = sql2 + " where platformAlias = ? and serverIp = ?"
		selectIds = []string{platformAlias,serverIp}
	}else{
		sql = sql2 + " where platformAlias = ?"
		selectIds = []string{platformAlias}
	}
	for key,value := range newplatDic{
		sql += " and "+ key + " =" + value
	}
	//if gameAlias == "lwtz"{
	dbname := beego.AppConfig.String(gameAlias+"mds")
	models.Init(dbname)
	o := orm.NewOrm()
	res,err := o.Raw(sql,selectIds).Exec()
	if err == nil{
		num,_ := res.RowsAffected()
		result["code"] = num
		result["msg"] = "更新成功！！！"
	}else{
		libs.NewLog().Error("faieds",err)
		result["msg"] = err
		result["code"] = "0"
	}

	u.Data["json"] = result
	u.ServeJSON()
	return



}

