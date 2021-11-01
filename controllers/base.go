package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type BaseController struct {
	beego.Controller
}

func (this *BaseController) init() {

}

func (this *BaseController) return_json(code int, msg string, data []map[])  {

}