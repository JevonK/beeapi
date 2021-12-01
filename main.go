package main

import (
	_ "beeapi/routers"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

func init()  {
	mysqlConfig, _ := beego.AppConfig.String("httpport")
	fmt.Print("1231231"+mysqlConfig+"\n")
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", mysqlConfig)
}

func main() {

	// 开启swagger api 文档
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
