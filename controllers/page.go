package controllers

import (
	"fmt"
	"strconv"

	//"fmt"
	"hello/models"

	"github.com/astaxie/beego"
	//config "github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/xml"
	//"github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/orm"
)

type PageController struct {
	beego.Controller
}

func (this *PageController) Online() {
	userId, _ := this.GetInt("userId")
	beego.Informational("userId值:")
	beego.Informational(userId)

	category, _ := this.GetInt("category")
	beego.Informational("category值:")
	beego.Informational(category)
	data := models.QueryCategoryTao(category, userId)
	fmt.Println(data)
	this.Data["json"] = data
	this.ServeJSON()

}

func (this *PageController) OnlineInit() {
	avatarUrl := this.GetString("avatarUrl")
	beego.Informational("avatarUrl值:")
	beego.Informational(avatarUrl)
	this.Data["AvatarUrl"] = avatarUrl

	nickName := this.GetString("nickName")
	beego.Informational("nickName值:")
	beego.Informational(nickName)
	this.Data["nickName"] = nickName

	userId, _ := this.GetInt("userId")
	beego.Informational("userId值:")
	beego.Informational(userId)

	category, _ := this.GetInt("category")
	beego.Informational("category值:" + strconv.Itoa(category))


	this.Data["AvatarUrl"] = avatarUrl
	this.Data["NickName"] = nickName
	this.TplName = "online.tpl"
}
