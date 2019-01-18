package libs

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
)


var (

	key1 []byte = []byte(beego.AppConfig.String("key1"))
	//key1 = beego.AppConfig.String("key1")
	key = beego.AppConfig.String("key")
)


//生成token
func GetToken(time int) string{
	claims := jwt.MapClaims{
		//NotBefore: int64(time.Now().Unix()),
		//NotBefore: int64(time),
		"iat": int64(time),
		//ExpiresAt: int64(time.Now().Unix() + 1000),
		//ExpiresAt: int64(time+1000),
		//Issuer:    key,
		"exp": int64(time + 1000),
		"iss": key,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key1)
	if err != nil {
		NewLog().Error("failed",err)
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
		NewLog().Error("failed",err)
		NewLog().Error("密钥验证错误，检查你的sign的值!!!!")
		fmt.Println("failed", err)
		return false
	}
	return true
}
