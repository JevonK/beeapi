package main

import (
	//_ "beeapi/routers"
	"fmt"
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

func init()  {
	mysqlConfig, _ := beego.AppConfig.String("httpport")
	fmt.Print("1231231"+mysqlConfig+"\n")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("nova_passport", "mysql", "root:123456@tcp(sandbox.test:3306)/nova_passport?charset=utf8mb4", 30)
	_, err := orm.GetDB("nova_passport")
	if err != nil {
		fmt.Println("get nova_passport DataBase")
	}
	o := orm.NewOrm()
	//var qs QuerySeter
	qs := o.QueryTable("banner")
	fmt.Println("get nova_passport DataBase: %s",qs)

	orm.RegisterModel(new(Banner))
}

func main() {
	//o := orm.NewOrm()
	//// read one
	//u := Banner{Id: 1}
	//err = o.Read(&u)
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	fmt.Println(u.id, u.pic)
	//}
	// 开启swagger api 文档
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
