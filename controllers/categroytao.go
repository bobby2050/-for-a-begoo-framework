package controllers

import (
	"fmt"
	// "hello/models"
	"strconv"

	"github.com/astaxie/beego"
	// config "github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/xml"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/orm"
)

type CategroyTaoController struct {
	beego.Controller
}

func (c *CategroyTaoController) Get() {
	// 获取到的code
	idString := c.Input().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("[err]:", err)
	}
	fmt.Println("[id]", id)

	c.Data["json"] = ""
	c.ServeJSON()
}
