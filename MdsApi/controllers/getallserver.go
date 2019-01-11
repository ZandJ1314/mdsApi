package controllers

import (
	"MdsApi/libs"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ServerController struct {
	beego.Controller
}

func (self * ServerController) Get(){
	result := make(map[string]interface{})
	dataurl := self.GetString("dataurl")
	time := self.GetString("time")
	sign := self.GetString("sign")
	gameId := self.GetString("gameId")
	systemId := self.GetString("systemId")
	libs.NewLog().Debug("使用的地址为%s,时间是%s",dataurl,time)
	libs.NewLog().Debug("传入的数据是gameId:%s,systemId:%s",gameId,systemId)
	token := GetToken(sign)
	fmt.Println(token)
	sql := "select serverId,platformId,platformAlias,gameDBname from mds_server where isDeleted = 0"
	plattest := make(map[string]interface{})
	var platarry []interface{}

	if CheckToken(token){
		if gameId == "28"{
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
				result["result"] = "false"
				result["msg"] = err
				libs.NewLog().Error("数据库执行错误！！！,请正确输入参数",err)
			}
		}else{
			o1 := orm.NewOrm()
			o1.Using("db2")
			var list []orm.ParamsList
			res,err := o1.Raw(sql).ValuesList(&list)
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
				result["result"] = "false"
				result["msg"] = err
				libs.NewLog().Error("数据库执行错误！！！,请正确输入参数",err)
			}
		}
	}else{
		libs.NewLog().Error("密钥验证错误，检查你的sign的值!!!!")
	}
	self.Data["json"] = result
	self.ServeJSON()
	return

}
