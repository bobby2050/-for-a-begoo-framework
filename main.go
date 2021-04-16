package main

import (
	"hello/models"
	_ "hello/routers"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
)

func init() {

	models.RegisterDB()
}
func main() {
	beego.Run()
}
