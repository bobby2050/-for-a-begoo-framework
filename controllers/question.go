package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	// config "github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/xml"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/orm"
)

type QuestionController struct {
	beego.Controller
}

func (c *QuestionController) Get() {
	// 获取到的type
	typeString := c.Input().Get("type")
	typeInt, _ := strconv.Atoi(typeString)
	fmt.Println("typeInt:", typeInt)

	// 获取到的lastQuestionId
	lastQuestionId, _ := c.GetInt("lastQuestionId")

	fmt.Println("lastQuestionId:", lastQuestionId)

	// 获取到的code
	categroyId,_ := c.GetInt("categroy")
	fmt.Println("categroy:", categroyId)

	// 每页记录数
	pagesize := 10
	fmt.Println("pagesize:", pagesize)

	// 页码
	perString := c.Input().Get("per")
	per, _ := strconv.Atoi(perString)
	fmt.Println("per:", per)

	userId, _ := c.GetInt("userId")

	fmt.Println("userId: ")
	fmt.Println(userId)

	if typeInt == 2 {
		fmt.Println("第二")
		//data := models.GetQuestionError(userId, per, pagesize, typeInt, lastQuestionId, categroyId)
		//c.Data["json"] = data

		c.TplName = "questionListError.tpl"

		//c.ServeJSON()
	} else {
		fmt.Println("正常")

		//data := models.GetQuestion(userId, categroyId, per, pagesize, typeInt, lastQuestionId)
		//c.Data["Data"] = data

		c.TplName = "questionList.tpl"
	}

}
