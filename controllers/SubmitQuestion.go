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

type SubmitQuestionController struct {
	beego.Controller
}

func (c *SubmitQuestionController) Post() {
	fmt.Println("-----------------------------------------------")
	// 获取到userId
	userId, _ := c.GetInt("userId")

	fmt.Println("userId: ", userId)

	// 获取到questionTaoId
	categroyTaoId, _ := c.GetInt("categroyTaoId")
	fmt.Println("categroyTaoId: ", categroyTaoId)

	// 获取到categroy2
	categroy2, _ := c.GetInt("categroy2")
	fmt.Println("categroy2: ", categroy2)

	// 获取到questionId
	questionIdString := c.Input().Get("questionId")
	questionId, _ := strconv.Atoi(questionIdString)
	fmt.Println("questionId: ", questionId)

	ok, _ := c.GetInt("ok")
	// 记录用户错误编号
	models.SubmitUserErrorQuestion(userId, questionId, categroyTaoId, ok, categroy2)

	// 记录最后的用户的id
	//models.AddUserLastQuestion(userId, questionId, categroyTaoId)

	c.ServeJSON()

}
