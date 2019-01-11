package controllers

import (
	"MdsApi/libs"
	_ "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type OperaTionQuery struct {
	beego.Controller
}

var (

	key []byte = []byte(beego.AppConfig.String("key"))
)

func (q *OperaTionQuery) Get() {
	result := make(map[string]interface{})
	platDic := make(map[string]string)
	newplatDic := make(map[string]string)
	//fmt.Println("testters")
	dataurl := q.GetString("dataurl")
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
	time := q.GetString("time")
	sign := q.GetString("sign")
	for key,value := range platDic{
		if strings.TrimSpace(value) != ""{
			newplatDic[key] = value
		}
	}
	libs.NewLog().Debug("使用的地址为%s,时间是%s",dataurl,time)
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

	token := GetToken(sign)
	fmt.Println(token)
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

	if CheckToken(token){
		//查询龙纹天尊数据库
		if gameAlias == "lwtz"{
			o := orm.NewOrm()
			if count == 1{
				var maps []orm.Params
				//selectIds := []string{platformAlias,serverId,serverIp,configId,combinedTo,isCombined}
				//sql := "select " + fields
				//sql = sql + " from mds_server where platformAlias = ? and serverId = ? and serverIp=? and configId=? and combinedTo=? and isCombined=?"
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
				//selectIds := []string{platformAlias,serverId,serverIp,configId,combinedTo,isCombined}
				//sql := "select " + fields
				//sql = " from mds_server where platformAlias = ? and serverId = ? and serverIp=? and configId=? and combinedTo=? and isCombined=?"
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

			//查询烈火风神数据库
		}else{
			o1 := orm.NewOrm()
			o1.Using("db2")
			if count == 1 {
				var maps []orm.Params
				//selectIds := []string{platformAlias,serverId,serverIp,configId,combinedTo,isCombined}
				//sql := "select " + fields
				//sql = " from mds_server where platformAlias = ? and serverId = ? and serverIp=? and configId=? and combinedTo=? and isCombined=?"
				res,err := o1.Raw(sql,selectIds).Values(&maps)
				fmt.Println(res,err)
				if err == nil && res > 0{
					result["code"] = "1"
					result["server_info"] = maps[0][fields]

				}else{
					result["code"] = "0"
					result["msg"] = err
					libs.NewLog().Error("数据库执行错误！！！,请正确输入参数",err)
				}
			}else{

				var list []orm.ParamsList
				//selectIds := []string{platformAlias,serverId,serverIp,configId,combinedTo,isCombined}
				//sql := "select " + fields
				//sql = " from mds_server where platformAlias = ? and serverId = ? and serverIp=? and configId=? and combinedTo=? and isCombined=?"
				res,err := o1.Raw(sql,selectIds).ValuesList(&list)
				if err == nil && res > 0{
					result["code"] = "1"
					result["server_info"] = list[0]
				}else{
					result["code"] = "0"
					result["msg"] = err
					libs.NewLog().Error("数据库执行错误！！！,请正确输入参数",err)
				}
			}
		}
	}else{
		libs.NewLog().Error("密钥验证错误，检查你的sign的值!!!!")
	}
	q.Data["json"] = result
	q.ServeJSON()
	return
}


//生成token
func GetToken(sign string) string{
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    sign,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(beego.AppConfig.String("key"))
	if err != nil {
		libs.NewLog().Error("failed",err)
		logs.Error(err)
		return ""
	}
	return ss
}


//校验token是否有效
func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		libs.NewLog().Error("failed",err)
		fmt.Println("failed", err)
		return false
	}
	return true
}

