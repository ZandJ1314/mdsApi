package controllers

import (
	"MdsApi/libs"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TestQuery struct {
	beego.Controller
}

var (

	key1 []byte = []byte(beego.AppConfig.String("key"))
)

func (t *TestQuery) Get()  {
	result := make(map[string]interface{})
	//gameAlias := "lwtz"
	fields := "serverName,javaDir"
	platformAlias := "yy_zsfy"
	serverId := "32000"
	serverIp := "39.107.83.59"
	configId := "1"
	sign := "zhangfan"
	token := GetToken1(sign)
	fmt.Println(token)
	if CheckToken1(token){
		o := orm.NewOrm()
		o.Using("db2")
		//var maps []orm.Params
		var list []orm.ParamsList
		selectIds := []string{platformAlias,serverId,serverIp,configId}
		sql2 := "select " + fields
		sql := sql2 + " from mds_server where platformAlias = ? and serverId = ? and serverIp=? and configId=?"
		//fmt.Println(sql)
		//res,err := o.Raw(sql,selectIds).Values(&maps)
		res,err := o.Raw(sql,selectIds).ValuesList(&list)
		fmt.Println(res,err)
		//fmt.Println(maps)
		fmt.Println(list)
		if err == nil && res > 0 {
			//fmt.Println("wwww")
			//fmt.Println(maps[0]["serverName"])
			fmt.Println(list[0][0])
			//result["server_info"] = maps[0][fields]
			result["server_info"] = list[0]
			result["code"] = "0"
		}else{
			fmt.Println("kongkongruye")
		}

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


