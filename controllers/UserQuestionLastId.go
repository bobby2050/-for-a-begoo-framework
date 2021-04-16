package controllers

import (
	"fmt"
	"hello/models"
	"strconv"

	"github.com/astaxie/beego"
	// config "github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/xml"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/orm"
)

type UserQuestionLastIdController struct {
	beego.Controller
}

func (c *UserQuestionLastIdController) Get() {

	// 获取到userId
	userIdString := c.Input().Get("userId")
	userId, _ := strconv.Atoi(userIdString)
	fmt.Println("userId: ", userId)

	// 获取到questionTaoId
	categroyId, _ := c.GetInt("categroyId")
	fmt.Println("categroyId: ", categroyId)

	// 获取到questionTaoId
	categroy2, _ := c.GetInt("categroy2")
	fmt.Println("categroy2: ", categroy2)

	// 记录最后用户的套题id
	data := models.GetUserLastQuestionId(userId, categroyId, categroy2)
	c.Data["json"] = data
	c.ServeJSON()

}
