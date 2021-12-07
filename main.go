package main

import (
	_ "beeapi/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	beego "github.com/beego/beego/v2/server/web"
)

type Banner struct {
	Id        int    `json:id`
	Title     string `json:title`
	Pic       string `json:pic`
	LinkUrl   string `json:link_url`
	Type      int    `json:type`
	POrder    string `json:p_order`
	IsShow    int    `json:is_show`
	CreatedAt int    `json:created_at`
	UpdatedAt int    `json:updated_at`
}

func init() {

}

func main() {
	start_service() // 启动的服务

	// 开启swagger api 文档
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func start_service()  {
	mysql_init()
}

func mysql_init()  {
	dsn, _ := beego.AppConfig.String("mysql::dsn")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
}
