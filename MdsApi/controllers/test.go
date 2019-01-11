package controllers

import (
	"MdsApi/libs"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type TestQuery struct {
	beego.Controller
}

var (

	key1 []byte = []byte(beego.AppConfig.String("key"))
)

func (t *TestQuery) Get()  {
	gamesq := t.GetString("name")
	if gamesq == ""{
		fmt.Println("hhhhh")
	}
	fmt.Println(gamesq)
	result := make(map[string]interface{})
	platDic := make(map[string]string)
	newplatDic := make(map[string]string)
	//gameAlias := "lwtz"
	fields := "serverName,javaDir"
	platformAlias := "yy_zsfy"
	libs.NewLog().Debug("操作的地址为%s,操作时间为%s",fields,platformAlias)
	serverId := "32000"
	platDic["serverId"] = serverId
	//serverIp := "39.107.83.59"
	//platDic["serverIp"] = serverIp
	configId := "1"
	platDic["configId"] = configId
	sign := "zhangfan"
	for key,value := range platDic{
		if strings.TrimSpace(value) != ""{
			newplatDic[key] = value
		}
	}

	token := GetToken1(sign)
	fmt.Println(token)
	plattest := make(map[string]interface{})
	if CheckToken1(token){
		o := orm.NewOrm()
		o.Using("db2")
		//sql := "select platformAlias,id,configId,gameAlias,gameId,platformName from mds_server group by platformAlias"
		sql := "select serverId,platformId,platformAlias from mds_server where isDeleted = 0"
		var list []orm.ParamsList
		var platarry []interface{}
		res,err := o.Raw(sql).ValuesList(&list)
		if err == nil && res>0{
			for i := 0;i<len(list);i++{
				//plattest["platformAlias"] = list[i][0]
				//plattest["id"] = list[i][1]
				//plattest["configId"] = list[i][2]
				//plattest["gameAlias"] = list[i][3]
				//plattest["gameId"] = list[i][4]
				//plattest["platformName"] = list[i][5]
				//platarry = append(platarry, plattest)
				plattest["serverId"] = list[i][0]
				plattest["platformId"] = list[i][1]
				plattest["platformAlias"] = list[i][2]
				platarry = append(platarry, plattest)
			}
			result["platformList"] = platarry
			fmt.Println(plattest)
		}
		//updatesql := "update mds_server set serverId = 33000 where platformAlias = yy_zsfy and serverId = 32000 and configId = 1;"
		//res,err := o.Raw(updatesql).Exec()
		//if err == nil{
		//	num,_ := res.RowsAffected()
		//	fmt.Println(num)
		//	result["code"] = num
		//}else{
		//	libs.NewLog().Error("faieds",err)
		//	result["msg"] = err
		//	result["code"] = "0"
		//}
		//var maps []orm.Params
		//var list []orm.ParamsList
		//var selectIds []string
		//
		//sql2 := "select " + fields
		//sql := sql2 + " from mds_server where platformAlias= ?"
		//for key,value := range newplatDic{
		//	sql += " and "+ key + " =" + value
		//	selectIds = []string{platformAlias}
		//}
		//fmt.Println(sql)
		////res,err := o.Raw(sql,selectIds).Values(&maps)
		//res,err := o.Raw(sql,selectIds).ValuesList(&list)
		//fmt.Println(res,err)
		////fmt.Println(maps)
		//fmt.Println(list)
		//if err == nil && res > 0 {
		//	//fmt.Println("wwww")
		//	//fmt.Println(maps[0]["serverName"])
		//	fmt.Println(list[0][0])
		//	//result["server_info"] = maps[0][fields]
		//	result["server_info"] = list[0]
		//	result["code"] = "0"
		//}else{
		//	fmt.Println("kongkongruye")
		//}

		t.Data["json"] = result
		t.ServeJSON()
		return
	}else{
		fmt.Println("failed check")
	}
	//combinedTo := "50001"
	//isCombined := "1"

}


//生成token
func GetToken1(sign string) string{
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    sign,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key1)
	if err != nil {
		//libs.NewLog().Error(err)
		libs.NewLog().Error("failed",err)
		//logs.Error(err)
		return ""
	}
	return ss
}


//校验token是否有效
func CheckToken1(token string) bool{
	_,err := jwt.Parse(token,func(*jwt.Token) (interface{},error){
		return key1,nil
	})
	if err != nil{
		libs.NewLog().Error("failed",err)
		fmt.Println("failed",err)
		return false
	}
	return true
}


