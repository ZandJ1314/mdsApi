package controllers

import (
	"MdsApi/libs"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AagentController struct {
	beego.Controller
}


func (a *AagentController) Get(){
	result := make(map[string]interface{})
	dataurl := a.GetString("dataurl")
	time := a.GetString("time")
	sign := a.GetString("sign")
	gameId := a.GetString("gameId")
	systemId := a.GetString("systemId")
	libs.NewLog().Debug("使用的地址为%s,时间是%s",dataurl,time)
	libs.NewLog().Debug("传入的数据是gameId:%s,systemId:%s",gameId,systemId)
	token := GetToken(sign)
	fmt.Println(token)
	sql := "select platformAlias,id,configId,gameAlias,gameId,platformName from mds_server group by platformAlias"
	plattest := make(map[string]interface{})
	var platarry []interface{}
	if CheckToken(token){
		if gameId == "28"{
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
		}else{
			o1 := orm.NewOrm()
			o1.Using("db2")
			var list []orm.ParamsList
			res,err := o1.Raw(sql).ValuesList(&list)
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
		}
	}else{
		libs.NewLog().Error("密钥验证错误，检查你的sign的值!!!!")
	}
	a.Data["json"] = result
	a.ServeJSON()
	return

}