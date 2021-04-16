package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})                                    // 接口
	beego.Router("/login", &controllers.UserController{})                               // 接口
	beego.Router("/getCategroy", &controllers.CategroyController{})                     // 接口
	beego.Router("/getCategroyTao", &controllers.CategroyTaoController{})               // 接口
	beego.Router("/getQuestion", &controllers.QuestionController{})                     // 网页
	beego.Router("/getQuestionFenye", &controllers.QuestionFenyeController{})           // 接口
	beego.Router("/submitQuestion", &controllers.SubmitQuestionController{})            // 接口 提交用户访问记录
	beego.Router("/reSubmitQuestion", &controllers.ReSubmitQuestionController{})        // 接口 再次提交用户访问记录
	beego.Router("/getUserQuestionLastId", &controllers.UserQuestionLastIdController{}) // 接口 获取用户访问记录
	beego.Router("/getUserQuestionError", &controllers.UserQuestionErrorController{})   // 接口 获取用户错误题记录

	beego.Router("/page", &controllers.PageController{}, "get:OnlineInit")
	beego.Router("/getData", &controllers.PageController{}, "get:Online") //  在线答题
	beego.Router("/wxbizdatacrypt", &controllers.UserController{}, "get:Wxbizdatacrypt") //  在线答题
	beego.Router("/register", &controllers.UserController{}, "get:Register") //  注册页面
	beego.Router("/reg", &controllers.UserController{}, "post:Reg") //  提交页面
	beego.Router("/saveNickName", &controllers.UserController{}, "post:SaveNickName") //  接口 授权后保存用户昵称

	beego.Router("/sendSms", &controllers.SmsController{}, "post:SendSms") //  接口 发送短信
	beego.Router("/submitUserInfo", &controllers.UserController{}, "post:SubmitUserInfo") //  接口 发送短信




}
