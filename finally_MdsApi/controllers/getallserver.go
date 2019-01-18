package controllers

import (
	"MdsApi/libs"
	"MdsApi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ServerController struct {
	BaseController
}

func (self * ServerController) Post(){
	result := make(map[string]interface{})
	gameId := self.GetString("gameId")
	systemId := self.GetString("systemId")
	libs.NewLog().Debug("传入的数据是gameId:%s,systemId:%s",gameId,systemId)
	sql := "select serverId,platformId,platformAlias,gameDBname from mds_server where isDeleted = 0"
	plattest := make(map[string]interface{})
	var platarry []interface{}

	dbname := beego.AppConfig.String(gameId+"mds")
	models.Init(dbname)
	o := orm.NewOrm()
	var list []orm.ParamsList
	res,err := o.Raw(sql).ValuesList(&list)
	if err == nil && res > 0{
		for i := 0;i<len(list);i++{
			plattest["serverId"] = list[i][0]
			plattest["platformId"] = list[i][1]
			plattest["platformAlias"] = list[i][2]
			plattest["gameDBName"] = list[i][3]
			platarry = append(platarry, plattest)
		}
		result["result"] = "true"
		result["serverListStr"] = platarry

	}else{
		libs.NewLog().Error("密钥验证错误，检查你的sign的值!!!!")
	}
	self.Data["json"] = result
	self.ServeJSON()
	return

}
