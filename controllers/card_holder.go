package controllers

import "beeapi/models"

// Operations about CardHolder
type CardHolderController struct {
	BaseController

}

func (this *CardHolderController) Get()  {
	list := models.GetCardHolder()
	this.Data["json"] = list
	this.return_json(200, "success", list)
}