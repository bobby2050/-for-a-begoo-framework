package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"hello/models"

	//config "github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/xml"
	//"github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/orm"
)

type IndexController struct {
	beego.Controller
}



func (this *IndexController) Index() {
	//avatarUrl := this.GetString("avatarUrl")
	users := models.GetUserList(1)
	fmt.Println(users)
	this.Data["Users"] = users
	this.TplName = "OrdinaryUser.tpl"
}

func (this *IndexController) Vip() {
	//avatarUrl := this.GetString("avatarUrl")
	users := models.GetUserList(2)
	fmt.Println(users)
	this.Data["Users"] = users
	this.TplName = "VipUser.tpl"
}


func (this *IndexController) MyOrder() {
	userId,_ := this.GetInt("userId")
	data := models.GetMyOrder(userId)
	fmt.Println(data)
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *IndexController) SubmitMyOrder() {
	userId,_ := this.GetInt("userId")
	categoryIds := this.GetString("categoryIds")
	models.SubmitMyOrder(userId, categoryIds)

	this.Data["json"] = ""
	this.ServeJSON()
}