package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
	"hello/util"
)

type SmsController struct {
	beego.Controller
}

func (this *SmsController) SendSms() {
	mobile := this.GetString("mobile")
	//beego.Informational("userId值:")
	beego.Informational("发送短信")

	code := util.GenValidateCode(6)

	beego.Informational(mobile)
	beego.Informational(code)

	models.SendSms(mobile, code)

	data := "ok"
	this.Data["json"] = data
	this.ServeJSON()

}

