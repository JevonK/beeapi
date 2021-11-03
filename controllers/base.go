package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

// Operations about object
type BaseController struct {
	beego.Controller
}

type returnData struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data map[interface{}]interface{} `json:"data"`
}

var (
	codeMsg = map[int]string{
		200: "请求成功",
		404: "非法路由",
		500: "服务器错误",
	}
)

func (this *BaseController) init() {

}

func (b *BaseController) return_json(code int, msg string, data interface{}) {
	logs.SetLogger(logs.AdapterFile, `{"color":true, "fileName": "logs/response.log"}`)
	if msg != "" {
		msg = codeMsg[code]
	}
	logs.Info("code: %d; msg: %s; data: %v ", code, msg, data)
	res := map[string] interface{}{"code": code, "msg": codeMsg[code], "data": data}
	b.Data["json"] = res
	b.ServeJSON()
}

func (b *BaseController) file_save_set (path string) *logs.BeeLogger {
	// 设置日志
	log := logs.NewLogger(10000)
	log.EnableFuncCallDepth(true)
	dateStr := time.Now().Format("2006-01-02")
	path = "logs/" + dateStr + "/" + path

	log.SetLogger(logs.AdapterFile, `{"color":true, "fileName": "` + path + `"}`)
	return log
}