package controllers

import (
	"fmt"
	"hello/models"

	"github.com/astaxie/beego"
	config "github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/xml"
	"github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/orm"
	"github.com/xlstudio/wxbizdatacrypt"
)

type UserController struct {
	beego.Controller
}

type Mystruct struct {
	User string
	Sex  string
	Age  int
	Code string
	Id   int
}

type WxStruct struct {
	Session_key string
	Openid      string
}

type WxPhoneNumber struct {
	PhoneNumber string
	PurePhoneNumber string
}
func (c *UserController) Get() {
	// 获取到的code
	code := c.Input().Get("code")
	fmt.Println("[code]:" + code)
	// 读取到的配置
	iniconf, err := config.NewConfig("ini", "conf/app.conf")

	if err != nil {
		fmt.Println(err)
	}
	appId := iniconf.String("appId")
	appSecret := iniconf.String("appSecret")
	fmt.Println("[appId]:" + appId)
	fmt.Println("[appSecret]:" + appSecret)

	// 请求微信拿到openid
	req := httplib.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + appSecret + "&js_code=" + code + "&grant_type=authorization_code")

	str, err := req.String()

	var wxdata WxStruct
	req.ToJSON(&wxdata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[请求微信拿到]:")
	fmt.Println(str)
	fmt.Println("Openid: " + wxdata.Openid)
	fmt.Println("Session_key: " + wxdata.Session_key)

	// var userData models.User
	// 创建一个 ormer 对象
	// o := orm.NewOrm()
	// user := new(models.User)
	// user.Name = "slene"

	// fmt.Println(o.Insert(user))
	id := models.AddUser(wxdata.Openid, wxdata.Session_key)

	// o.Insert(&user)

	// 返回数据
	data := &Mystruct{User: "bbb", Sex: "true", Age: 12, Code: code, Id: id}
	c.Data["json"] = data

	c.ServeJSON()
}

func (this *UserController) LookUser() {
	this.Data["json"] = "hello"

	this.ServeJSON()
}

func (this *UserController) Wxbizdatacrypt() {

	// 读取到的配置
	userId,_ := this.GetInt("userId")
	sessionKey := models.GetUserSessionKey(userId);
	encryptedData := this.GetString("encryptedData")
	iv := this.GetString("iv")


	appID := beego.AppConfig.String("appId")

	pc := wxbizdatacrypt.WxBizDataCrypt{AppId: appID, SessionKey: sessionKey}
	result, err := pc.Decrypt(encryptedData, iv, true) //第三个参数解释： 需要返回 JSON 数据类型时 使用 true, 需要返回 map 数据类型时 使用 false
	if err != nil {
		fmt.Println("微信解密后的失败：")
		fmt.Println(err)
	} else {
		fmt.Println("微信解密成功，返回的数据：")
		fmt.Println(result)
	}

	this.Data["json"] = result

	this.ServeJSON()
}

func (this *UserController) Register() {
	userId, _ := this.GetInt("userId")
	re := models.CheckUserData(userId)
	if re == "fail" {
		this.TplName = "no.tpl"
	} else {
		this.TplName = "register.tpl"
	}

}

func (this *UserController) Reg() {
	username := this.GetString("username")
	position := this.GetString("position")
	//mobileNum := this.GetString("mobileNum")
	userId,_ := this.GetInt("userId")

	models.AppendUserInfo(username, position, userId)
	result := username
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *UserController) SaveNickName() {
	nick_name := this.GetString("nick_name")
	user_id,_ := this.GetInt("user_id")
	models.SaveUserInfo(nick_name, user_id)
	this.Data["json"] = ""
	this.ServeJSON()
}

func (this *UserController) SubmitUserInfo() {
	username := this.GetString("username")
	position := this.GetString("position")
	mobileNum := this.GetString("mobileNum")
	reCode := this.GetString("reCode")
	userId,_ := this.GetInt("userId")

	beego.Informational(username)
	beego.Informational(position)
	beego.Informational(mobileNum)
	beego.Informational(reCode)
	beego.Informational(userId)
	flag := models.SubmitUserInfo(userId, reCode, username, position, mobileNum)
	this.Data["json"] = flag
	this.ServeJSON()

}

