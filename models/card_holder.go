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

func GetCardHolder() (holder *CardHolder) {
	var o = orm.NewOrm()

	cardHolder := CardHolder{Id: 1}
	err := o.Read(&cardHolder)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键 ")
	} else {
		fmt.Println(cardHolder.Id, cardHolder.SkuNo)
	}
	return holder
}
