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

type QuestionFenyeController struct {
	beego.Controller
}

func (c *QuestionFenyeController) Get() {

	// 获取到的type
	typeString := c.Input().Get("type")
	typeInt, _ := strconv.Atoi(typeString)
	fmt.Println("typeInt:", typeInt)

	// 获取到的lastQuestionId
	lastQuestionIdString := c.Input().Get("lastQuestionId")
	lastQuestionId, _ := strconv.Atoi(lastQuestionIdString)
	fmt.Println("lastQuestionId:", lastQuestionId)

	// 获取到的code
	categroyId, _ := c.GetInt("categroy")

	fmt.Println("categroyId:", categroyId)

	//每页记录数
	pagesize := 10
	fmt.Println("pagesize:", pagesize)

	// 页码
	perString := c.Input().Get("per")
	per, _ := strconv.Atoi(perString)
	fmt.Println("per:", per)

	userIdString := c.Input().Get("userId")
	userId, _ := strconv.Atoi(userIdString)

	fmt.Println("userId: ")
	fmt.Println(userId)

	categroy, _ := c.GetInt("category")
	fmt.Println("categroy:", categroy)


	if typeInt == 2 {

		data := models.GetQuestionError(userId, per, pagesize, typeInt, lastQuestionId, categroy)
		c.Data["json"] = data
		c.ServeJSON()
	} else {
		data := models.GetQuestion(userId, categroyId, per, pagesize, typeInt, lastQuestionId, categroy)
		c.Data["json"] = data
		c.ServeJSON()
	}

}
