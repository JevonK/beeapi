package controllers

import (
	"beeapi/models"
	"fmt"
)

// Operations about CardHolders
type CardHolderController struct {
	BaseController
}

// @Title Get
// @Description get all the staticblock by key
// @Param	body		body 	models.CardHolder	true		"body for CardHolder content"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [get]
func (this *CardHolderController) Get() {
	fmt.Print("12313")
	this.Data["json"] = "12331"
	this.ServeJSON()
	//list := models.GetCardHolder()
	//this.Data["json"] = list
	//this.return_json(200, "success", list)
}

// @Title Post
// @Description get all the staticblock by key
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [post]
func (this *CardHolderController) Post() {
	list := models.GetCardHolder()
	this.return_json(200, "success", list)
}

// @Title get
// @Description get all the staticblock by key
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /asd [get]
func (this *CardHolderController) Asd() {
	fmt.Print("12313")
	this.Data["json"] = "12331"
	this.ServeJSON()
	//list := models.GetCardHolder()
	//this.Data["json"] = list
	//this.return_json(200, "success", list)
}