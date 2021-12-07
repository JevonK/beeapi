package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

//import "github.com/astaxie/beego/orm"

type CardHolder struct {
	Id         int    `json:"id"`
	SupplierId int    `json:"supplier_id"`
	SkuNo      string `json:"sku_no"`
	CardNumber string `json:"card_number"`
	CreatedAt  int    `json:"created_at"`
	IsIssue    int    `json:"is_issue"`
	IsUsed     int    `json:"is_used"`
	OrderId    string `json:"order_id"`
	IssueTime  int    `json:"issue_time"`
	UseTime    int    `json:"use_time"`
	IsValid    int    `json:"is_valid"`
}


func init()  {
	orm.RegisterModel(new(CardHolder))
}

func add()  {

}

func GetCardHolder() (holders interface{}) {
	var o = orm.NewOrm()
	var card_holder_list []*CardHolder
	_, err := o.QueryTable("CardHolder").All(&card_holder_list)

	if err != nil {
		fmt.Println("查询不到")
	}
	return card_holder_list
}
