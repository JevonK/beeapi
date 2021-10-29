package main

import (
	_ "beeapi/routers"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	log := logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"logs/test.log"}`)
	//val, _ := beego.AppConfig.String("appname")
	log.Info("load config name is")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
