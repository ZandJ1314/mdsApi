package controllers

import (
	"MdsApi/libs"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type UpdateController struct {
	beego.Controller
}

func (u *UpdateController) Get(){
	result := make(map[string]interface{})
	platDic := make(map[string]string)
	newplatDic := make(map[string]string)
	//fmt.Println("testters")
	dataurl := u.GetString("dataurl")
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
	time := u.GetString("time")
	sign := u.GetString("sign")
	for key,value := range platDic{
		if strings.TrimSpace(value) != ""{
			newplatDic[key] = value
		}
	}
	libs.NewLog().Debug("使用的地址为%s,时间是%s",dataurl,time)

	token := GetToken(sign)
	fmt.Println(token)
	var logstr string
	logstr = "传入后端的数据有"
	for key,value := range newplatDic{
		logstr += ","+ key + ":" + value
	}
	if serverIp != ""{
		logstr += ",fields:%s,platformAlias:%s,gameAlias:%s,serverIp:%s"
		libs.NewLog().Debug(logstr,fields,platformAlias,gameAlias,serverIp)
	}else{
		logstr += ",fields:%s,platformAlias:%s,gameAlias:%s,serverIp:%s"
		libs.NewLog().Debug(logstr,fields,platformAlias,gameAlias)
	}

	var selectIds []string
	var sql string
	sql2 := "update mds_server set "

	if CheckToken(token){
		//龙纹天尊
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
		if gameAlias == "lwtz"{
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
		}else{
			o1 := orm.NewOrm()
			o1.Using("db2")
			res,err := o1.Raw(sql,selectIds).Exec()
			if err == nil{
				num,_ := res.RowsAffected()
				result["code"] = num
			}else{
				libs.NewLog().Error("faieds",err)
				result["msg"] = err
				result["code"] = "0"
			}
		}
	}else{

		libs.NewLog().Error("密钥验证错误，检查你的sign的值!!!!")
	}

	u.Data["json"] = result
	u.ServeJSON()
	return



}

