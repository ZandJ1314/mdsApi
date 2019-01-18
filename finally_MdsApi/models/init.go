package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/url"
	_ "github.com/go-sql-driver/mysql"
)

func Init(dbname string){
	orm.RegisterDriver("mysql",orm.DRMySQL)
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	timezone := beego.AppConfig.String("db.timezone")
	mydbname := dbname

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + mydbname + "?charset=utf8"
	fmt.Println(dsn)
	//dsn2 := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + "zsfy_mds" + "?charset=utf8"
	//fmt.Println(dsn2)
	if timezone != ""{
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
		//dsn2 = dsn2 + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default","mysql",dsn)
	//orm.RegisterDataBase("db2","mysql",dsn2)
	fmt.Println("连接数据库成功")
	if beego.AppConfig.String("runmode") == "dev"{
		orm.Debug = true
	}
}

func TableName(name string) string  {
	return "mds_" + name
}