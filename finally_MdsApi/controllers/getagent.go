package controllers

import (
	"MdsApi/libs"
	"MdsApi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AagentController struct {
	BaseController
}


func (a *AagentController) Post(){
	result := make(map[string]interface{})

	gameId := a.GetString("gameId")
	systemId := a.GetString("systemId")
	libs.NewLog().Debug("传入的数据是gameId:%s,systemId:%s",gameId,systemId)

	sql := "select platformAlias,id,configId,gameAlias,gameId,platformName from mds_server group by platformAlias"
	plattest := make(map[string]interface{})
	var platarry []interface{}
	dbname := beego.AppConfig.String(gameId+"mds")
	models.Init(dbname)
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql).ValuesList(&list)
	if err == nil && res > 0{
		for i := 0;i<len(list);i++{
			plattest["platformAlias"] = list[i][0]
			plattest["id"] = list[i][1]
			plattest["configId"] = list[i][2]
			plattest["gameAlias"] = list[i][3]
			plattest["gameId"] = list[i][4]
			plattest["platformName"] = list[i][5]
			platarry = append(platarry, plattest)
		}
		result["result"] = "true"
		result["platformList"] = platarry
	}else{
		result["result"] = "false"
		result["msg"] = err
		libs.NewLog().Error("数据库执行错误！！！,请正确输入参数",err)
	}

	a.Data["json"] = result
	a.ServeJSON()
	return

}