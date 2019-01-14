package controllers

import (
	"MdsApi/libs"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type InsertMdsController struct {
	beego.Controller
}


func (i *InsertMdsController) Post(){
	result := make(map[string]interface{})
	jsontomap := make(map[string]interface{})
	dataurl := i.GetString("dataurl")
	time := i.GetString("time")
	sign := i.GetString("sign")
	gameId := i.GetString("gameId")
	systemId := i.GetString("systemId")
	serverListStr := i.GetString("serverListStr")
	libs.NewLog().Debug("使用的地址为%s,时间是%s",dataurl,time)
	libs.NewLog().Debug("传入的数据是gameId:%s,systemId:%s,serverListStr:%s",gameId,systemId,serverListStr)
	token := GetToken(sign)
	fmt.Println(token)

	if CheckToken(token){
		//json字符串转换为map
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

		if gameId == "28"{

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
		}else{

			o1 := orm.NewOrm()
			o1.Using("db2")
			res,err := o1.Raw(sql,insertmds).Exec()
			if err == nil{
				num,_ := res.RowsAffected()
				result["code"] = num
				result["msg"] = "插入数据成功！！！"
			}else{
				libs.NewLog().Error("faieds",err)
				result["msg"] = err
				result["code"] = "0"
			}
		}
	}else{
		libs.NewLog().Error("密钥验证错误，检查你的sign的值!!!!")
	}

	i.Data["json"] = result
	i.ServeJSON()
	return


}