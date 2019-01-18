package controllers

import (
	"MdsApi/libs"
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
}


func (self *BaseController) Prepare(){
	//获取前端传进来的url
	dataurl := self.Ctx.Request.URL
	result := make(map[string]interface{})
	localkey := beego.AppConfig.String("key")
	time := self.GetString("time")
	libs.NewLog().Debug("使用的地址为%s,时间戳是%s",dataurl,time)
	//fmt.Println(dataurl,time)
	rettime := strings.TrimSpace(time)
	//fmt.Println(rettime,localkey)
	securSign := libs.Md5([]byte(localkey+rettime))

	if securSign != ""{
		if self.auth(securSign){
			return
		}else{
			clientip := self.getClientIp()
			result["code"] = "0"
			result["msg"] = "密钥验证错误,请确认后重新执行,访问的Ip地址为:"+ clientip
			libs.NewLog().Error("密钥验证错误,请确认后重新执行,访问的ip地址为:%s",clientip)
		}
	}
	self.Data["json"] = result
	self.ServeJSON()
	return
}

func (self *BaseController) auth(str string) bool {
	sign := self.GetString("sign")
	sign = strings.TrimSpace(sign)
	if str != sign{
		return false
	}
	return true
}

//获取访问用户ip地址
func (self *BaseController) getClientIp() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}