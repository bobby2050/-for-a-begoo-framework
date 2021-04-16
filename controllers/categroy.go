package controllers

import (
	// "fmt"
	"hello/models"

	"github.com/astaxie/beego"
	// config "github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/xml"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/orm"
)

type CategroyController struct {
	beego.Controller
}

func (this *CategroyController) Get() {
	user_id, _ := this.GetInt("user_id")
	category := models.QueryCategory(user_id)
	this.Data["json"] = category
	this.ServeJSON()
}
