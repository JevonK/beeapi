package main

import (
	_ "beeapi/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// 开启swagger api 文档
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
