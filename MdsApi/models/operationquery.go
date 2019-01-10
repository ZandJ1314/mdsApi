package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/url"
	_ "github.com/go-sql-driver/mysql"
)

func Init(){
	orm.RegisterDriver("mysql",orm.DRMySQL)
	dbhost := "rm-2ze8syz2nh1jtm6tdmo.mysql.rds.aliyuncs.com"
	dbport := "3306"
	dbuser := "chaxun"
	dbpassword := "MS6IVvC9XFdQMi1gMwxK"
	timezone := "Asia/Shanghai"

	dsn1 := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + "lwtz_mds" + "?charset=utf8"
	fmt.Println(dsn1)
	dsn2 := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + "zsfy_mds" + "?charset=utf8"
	fmt.Println(dsn2)
	if timezone != ""{
		dsn1 = dsn1 + "&loc=" + url.QueryEscape(timezone)
		dsn2 = dsn2 + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default","mysql",dsn1)
	orm.RegisterDataBase("db2","mysql",dsn2)
	fmt.Println("连接数据库成功")
	if beego.AppConfig.String("runmode") == "dev"{
		orm.Debug = true
	}
}

func TableName(name string) string  {
	return "mds_" + name
}