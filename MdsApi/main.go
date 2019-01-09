package main

import (
	"MdsApi/models"
	_ "MdsApi/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
