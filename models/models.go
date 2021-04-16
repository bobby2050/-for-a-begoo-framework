package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"hello/util"
	"math"
	"strconv"
	"strings"

	// "golang.org/x/tools"

	// "strings"

	// "strings"
	"time"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql" // <-
)

// 用户(定义表)
type User struct {
	Id         int
	Sessionkey string
	Openid     string
	MobileNum  string    `orm:"null"` // 手机号
	UserType   int       `orm:"null"` // 1:普通用户 2：vip用户
	NickName   string    `orm:"null"` // 微信昵称
	UserName   string    `orm:"null"` // 姓名
	Position   string    `orm:"null"` // 岗位
	Created    time.Time `orm:"null"`
	Updated    time.Time `orm:"null"`
}

// 证书题库分类 (定义表)
type Category struct {
	Id      int
	Name    string
	Img     string
	Created time.Time `orm:"null"`
	Updated time.Time `orm:"null"`
}

// 用户开通证书题库(定义表)
type UserCategory struct {
	Id         int
	UserId     int       `orm:"null"` // 用户编号
	CategoryId int       `orm:"null"` // 证书题库
	Created    time.Time `orm:"null"`
	Updated    time.Time `orm:"null"`
}

// 证书题库分类套次 (定义表)
type CategoryTao struct {
	Id            int
	Name          string
	Ranking       int
	TotalQuestion int
	Category      int
	AnswerType    int // 单选还是多选
	Created       time.Time `orm:"null"`
	Updated       time.Time `orm:"null"`
}

//  题 (定义表)
type Question struct {
	Id             int
	Num            int                        // 题号
	Title          string                     // 题目
	Img            string                     // 图片
	AnswerAnalysis string `orm:"type(text) "` // 答案解析
	CategoryTao    int                        // 证书题库分类套次
	Created        time.Time `orm:"null"`
	Updated        time.Time `orm:"null"`
}

//  题选项  (定义表)
type Answer struct {
	Id           int
	Num          int    // 题号
	OptionName   string // 选项
	OptionAnswer int    // 选项是否正确
	CategoryTao  int    // 题库分类套次
	OrderNum     int    // 排序
	Created      time.Time `orm:"null"`
	Updated      time.Time `orm:"null"`
}

//  用户答题错误记录 (定义表)
type UserErrorQuestion struct {
	Id            int
	QuestionId    int // 题编号
	UserId        int // 用户id
	Status        int // 答题是否错误 1错误 2正确
	CategoryTaoId int // 证书题库分类套次
	Categroy      int //职称编号
	Created       time.Time `orm:"null"`
	Updated       time.Time `orm:"null"`
}

//  短信 (定义表)
type Sms struct {
	Id      int
	Mobile  string // 手机号
	Code    string // 短信码
	Created time.Time `orm:"null"`
	Updated time.Time `orm:"null"`
}

// 返回题库分类
type ReCategory struct {
	Id     int
	Name   string
	Img    string
	IsAuth int // 是否授权
}

// 返回题
type QuestionList struct {
	ReQuestion      []ReQuestion
	CategoryTaoName string
	Total1          int
	Total           float64
	UserId          int
	CategoryTaoId   int
	TotalPage       float64
}

// 返回题
type ReQuestion struct {
	Id             int                        // 编号
	Num            int                        // 题号
	Title          string                     // 题目
	Img            string                     // 图片
	AnswerAnalysis string `orm:"type(text) "` // 答案解析
	Answer         []Answer
	OptionType     int // 1 单选、2多选
}

// 返回已经和当前题库
type ReCategoryTao struct {
	Id             int
	Name           string
	Ranking        int
	TotalQuestion  int
	Category       int
	AnswerType     int // 单选还是多选
	IsReady        bool
	LastQuestionId int // 用户答题最后的题编号
}

// 返回已经
type ListReCategoryTao struct {
	ReCategoryTao  []ReCategoryTao
	LastQuestionId int
	MobileNum      string
	IsAuth         int // 是否订阅
}

func RegisterDB() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Category), new(CategoryTao), new(Question), new(Answer), new(UserErrorQuestion), new(UserCategory), new(Sms))

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:"+beego.AppConfig.String("mysqlpass")+"@tcp(localhost:"+beego.AppConfig.String("mysqlport")+")/cw?charset=utf8mb4&parseTime=true&loc=Local")
	orm.RunSyncdb("default", false, true)
	orm.Debug = true

	// 数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)

	// 数据库的最大数据库连接
	orm.SetMaxOpenConns("default", 30)
}

// 获取题库分类
func QueryCategoryTao(id int, userId int) ListReCategoryTao {
	var listReCategoryTao ListReCategoryTao
	var reCategoryTao []ReCategoryTao
	o := orm.NewOrm()
	// category := make([]*Category, 0)
	fmt.Println("[获取分类的id是] ： ", id)
	var categoryTao []CategoryTao
	num, err := o.Raw("SELECT id, name,ranking,total_question FROM category_tao where category=?", id).QueryRows(&categoryTao)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
	fmt.Println(userId)

	for _, n := range categoryTao {
		fmt.Println("当前：" + strconv.Itoa(n.Id))

		fmt.Println("无")
		nRow := ReCategoryTao{Id: n.Id, Name: n.Name, Ranking: n.Ranking, TotalQuestion: n.TotalQuestion}
		reCategoryTao = append(reCategoryTao, nRow)

	}
	listReCategoryTao.ReCategoryTao = reCategoryTao

	// 用户手机号
	var user User
	o.Raw("SELECT mobile_num from user where id=?", userId).QueryRow(&user)

	listReCategoryTao.MobileNum = user.MobileNum

	// 该套题是否订阅
	var userCategory UserCategory
	o.Raw("SELECT id from user_category where user_id=? and category_id= ? ", userId, id).QueryRow(&userCategory)
	fmt.Println("userCategory:")
	fmt.Println(userCategory)
	var isAuth int = 2
	if userCategory.Id > 0 {
		isAuth = 1
	}
	listReCategoryTao.IsAuth = isAuth; // 是否订阅 1 已经订阅，2未订阅

	return listReCategoryTao

}

// 每页显示10记录题
// typeInt： 1: 普通套题 2 错误集
func GetQuestion(userId int, categroyId int, per int, pagesize int, typeInt int, lastQuestionId int, categroy int) QuestionList {
	fmt.Println("------------------------GetQuestion-------------------")
	var questionList QuestionList
	o := orm.NewOrm()
	offset := 0

	// 拿到用户访问的最大记录，数据从最大记录开始

	fmt.Println("[per]:", per)
	if per == 0 || per == 1 {
		offset = 0
	} else if per > 1 {
		offset = (per - 1) * pagesize
	}

	// 拿到题库套名
	var categoryTao CategoryTao
	err := o.Raw("SELECT name  FROM category_tao where id=?", categroyId).QueryRow(&categoryTao)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[categoryTao]:")

	// 拿到该套错误的id
	var userErrorQuestion []UserErrorQuestion
	o.Raw("SELECT question_id FROM user_error_question WHERE user_id =? and category_tao_id=? ", userId, categroyId).QueryRows(&userErrorQuestion)
	nums := []int{}
	for _, v := range userErrorQuestion {
		nums = append(nums, v.QuestionId)
	}
	fmt.Println("nums:")
	fmt.Println(nums)
	// 从question库拿出题
	var question []Question
	// num, err := o.Raw("SELECT id,num,title,img ,answer_analysis FROM question WHERE category_tao=? and id not in(?) ORDER BY num limit ?,? ", categroyId, nums, offset, pagesize).QueryRows(&question)
	fmt.Println("=============")
	var totalData int

	// 拿去所有数据
	var str string
	for i := 0; i < len(nums); i++ {
		str += ",?"
	}

	str = strings.Trim(str, ",")
	fmt.Println("[in]:")
	fmt.Println(nums)
	fmt.Println(str)

	// 拿去所有数据(除过已经已答题)
	if str == "" {
		o.Raw("SELECT id,num,title,img ,answer_analysis FROM question WHERE category_tao=?  ", categroyId).QueryRows(&question)
		totalData = len(question)
	} else {
		o.Raw("SELECT id,num,title,img ,answer_analysis FROM question WHERE category_tao=? and id not in("+str+") ", categroyId, nums).QueryRows(&question)
		totalData = len(question)
	}
	fmt.Println("[per]:" + strconv.Itoa(per))
	tmpTotalData := float64(totalData) / float64(pagesize)
	fmt.Println("一共多数数据：")
	tmpTotalData = math.Ceil(tmpTotalData)
	fmt.Println(tmpTotalData)

	//fmt.Println("一共多数数据：" + strconv.f(tmpTotalData))
	fmt.Println("len(nums): " + strconv.Itoa(len(nums)))
	if len(nums) == 0 || lastQuestionId == 0 {
		fmt.Println("不存在lastid")
		num, _ := o.Raw("SELECT id,num,title,img ,answer_analysis FROM question WHERE category_tao=?  ORDER BY num limit ?,? ", categroyId, offset, pagesize).QueryRows(&question)
		fmt.Println(num)
		if num == 0 {
			return QuestionList{CategoryTaoName: categoryTao.Name, Total: 0, UserId: userId}
		}
	} else {
		fmt.Println("存在lastid")
		var str string
		for i := 0; i < len(nums); i++ {
			str += ",?"
		}

		str = strings.Trim(str, ",")
		fmt.Println("[in]:")
		fmt.Println(nums)
		fmt.Println(str)

		num, _ := o.Raw("SELECT id,num,title,img ,answer_analysis FROM question WHERE category_tao=? and id not in("+str+") ORDER BY num limit ?,? ", categroyId, nums, offset, pagesize).QueryRows(&question)
		fmt.Println(num)
		if num == 0 {
			return QuestionList{CategoryTaoName: categoryTao.Name, Total: 0, UserId: userId}
		}
	}

	fmt.Println("=============")

	// 拿到Num字段
	ids := []int{}
	for _, v := range question {
		//fmt.Println(v.Num)
		// inData += strconv.Itoa(v.Num) + ","
		ids = append(ids, v.Num)
	}
	// s := strings.TrimRight(inData, ",")
	fmt.Println("[s]:", ids)
	var answer []Answer

	o.QueryTable("answer").Filter("category_tao", categroyId).Filter("num__in", ids).All(&answer)
	//o.Raw("SELECT num,option_name,option_answer,order_num FROM answer WHERE category_tao=? and num IN (?, ?)", id, ids).QueryRows(&answer)

	// 最终返回
	var reQuestion []ReQuestion
	for _, v := range question {
		num := v.Num
		var newAnswer []Answer
		var optionType int
		for _, vv := range answer {
			if num == vv.Num {
				newAnswerRow := Answer{Num: vv.Num, OptionName: vv.OptionName, OptionAnswer: vv.OptionAnswer, OrderNum: vv.OrderNum}
				newAnswer = append(newAnswer, newAnswerRow)
			}

		}
		for _, vvv := range newAnswer {
			if vvv.OptionAnswer == 1 {
				optionType++
			}
		}
		newStruct := ReQuestion{Id: v.Id, Num: v.Num, Title: v.Title, Img: v.Img, AnswerAnalysis: v.AnswerAnalysis, Answer: newAnswer, OptionType: optionType}
		reQuestion = append(reQuestion, newStruct)

	}
	fmt.Println("----------------")
	fmt.Println(len(reQuestion))
	fmt.Println("----------------")

	questionList = QuestionList{ReQuestion: reQuestion, CategoryTaoName: categoryTao.Name, Total: tmpTotalData, UserId: userId}
	return questionList
}

// 获取题库分类
func QueryCategory(user_id int) []ReCategory {
	o := orm.NewOrm()
	// category := make([]*Category, 0)
	var userCategory []UserCategory
	o.Raw("SELECT category_id FROM user_category where user_id=?", user_id).QueryRows(&userCategory)

	var category []ReCategory
	num, err := o.Raw("SELECT id, name, img FROM category").QueryRows(&category)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range category {

		category[i].IsAuth = tmpFind(v.Id, userCategory)
	}

	fmt.Println(num)
	return category

}

func tmpFind(id int, userCategory []UserCategory) int {
	fmt.Println("当前id：" + strconv.Itoa(id))
	for _, v := range userCategory {
		if id == v.CategoryId {
			return 1
			break
		}
	}
	return 0
}

//  查询用户openid记录是否存在
func queryUser(openid string) int {
	var id int
	fmt.Println("查询")
	var user User
	o := orm.NewOrm()

	err := o.QueryTable("user").Filter("openid", openid).One(&user, "Id")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		id = 0
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		id = 0
	} else {
		fmt.Println("找到 :" + strconv.Itoa(user.Id))
		id = user.Id
	}
	return id
}

// 新增用户
func AddUser(openid string, session_key string) int {
	var id int = queryUser(openid)
	if id > 0 {
		fmt.Println("更新数据")
		user := User{Id: id}
		user.Sessionkey = session_key
		user.Updated = time.Now()
		o := orm.NewOrm()
		o.Using("default")

		user.Sessionkey = session_key

		user.Updated = time.Now()
		o.Update(&user, "Sessionkey", "Updated")
	} else {
		fmt.Println("新增数据")
		var user User
		o := orm.NewOrm()
		o.Using("default")

		user.Openid = openid
		user.Sessionkey = session_key
		user.UserType = 1
		user.Created = time.Now()
		user.Updated = time.Now()
		o.Insert(&user)
		id = user.Id
		fmt.Println("id:" + strconv.Itoa(id))
	}
	return id

}

// 保存用户的昵称
func SaveUserInfo(nick_name string, user_id int) {

	user := User{Id: user_id}
	user.NickName = nick_name
	user.Updated = time.Now()

	o := orm.NewOrm()
	o.Using("default")

	user.Updated = time.Now()
	o.Update(&user, "NickName", "Updated")
}

// 追加用户信息
func AppendUserInfo(username string, position string, userId int) {
	var user User
	o := orm.NewOrm()
	o.Raw("SELECT id FROM user_error_question WHERE user_id = ?", userId).QueryRow(&user)
	user.UserName = username
	user.Position = position
	//user.MobileNum = mobileNum

	o.Update(&user, "UserName", "Position")

}

// 记录用户错误编号
func SubmitUserErrorQuestion(userId int, questionId int, categroyTaoId int, ok int, categroy2 int) {
	fmt.Println("----------------------------------------")
	fmt.Println(questionId)

	o := orm.NewOrm()
	o.Using("default")
	var userErrorQuestion UserErrorQuestion

	o.Raw("SELECT id FROM user_error_question WHERE user_id = ? and question_id=? and category_tao_id=? and categroy = ? limit 1", userId, questionId, categroyTaoId, categroy2).QueryRow(&userErrorQuestion)

	var id = userErrorQuestion.Id
	fmt.Println(id)
	if id == 0 {
		fmt.Println("不存在，新增")
		var userErrorQuestion UserErrorQuestion
		userErrorQuestion.CategoryTaoId = categroyTaoId
		userErrorQuestion.QuestionId = questionId
		userErrorQuestion.UserId = userId
		userErrorQuestion.Created = time.Now()
		userErrorQuestion.Updated = time.Now()
		userErrorQuestion.Status = ok // 答题状态
		userErrorQuestion.Categroy = categroy2
		o.Insert(&userErrorQuestion)
	} else {
		fmt.Println("存在，修改")
		var userErrorQuestion UserErrorQuestion
		userErrorQuestion.Id = id
		userErrorQuestion.Updated = time.Now()
		userErrorQuestion.Status = ok // 答题状态
		userErrorQuestion.Categroy = categroy2
		o.Update(&userErrorQuestion, "Updated", "Status", "Categroy")
	}

	// 把本套id写入记录
	//var userLastCategoryTao UserLastCategoryTao
	//o.Raw("SELECT id  FROM user_last_category_tao WHERE query_category_tao_id = ? and user_id = ?", categroyTaoId, userId).QueryRow(&userLastCategoryTao)
	//fmt.Println("是否存在：")
	//fmt.Println(userLastCategoryTao)
	//if userLastCategoryTao.Id == 0 {
	//	var userLastCategoryTao UserLastCategoryTao
	//	userLastCategoryTao.QueryCategoryTaoId = categroyTaoId
	//	userLastCategoryTao.UserId = userId
	//	userLastCategoryTao.Created = time.Now()
	//	userLastCategoryTao.Updated = time.Now()
	//	o.Insert(&userLastCategoryTao)
	//}

	// 处理是否换下一套

	// 拿到question表TaoId的最大id
	var question []Question
	o.Raw("SELECT id  FROM question WHERE category_tao = ?", categroyTaoId).QueryRows(&question)
	questonNum := len(question)
	fmt.Println("questonNum:")
	fmt.Println(questonNum)

	// 两个id相等的话，user_error_question表
	var userErrorQuestion2 []UserErrorQuestion
	o.Raw("SELECT id  FROM user_error_question WHERE category_tao_id = ? and user_id= ? and status=2", categroyTaoId, userId).QueryRows(&userErrorQuestion2)
	userErrorQuestion2Num := len(userErrorQuestion2)
	fmt.Println("userErrorQuestion2Num:")
	fmt.Println(userErrorQuestion2Num)

	//if userErrorQuestion2Num == questonNum {
	//	fmt.Println("相等，无错误")
	//
	//	var categoryTao CategoryTao
	//	o.Raw("SELECT category FROM category_tao where id=?", categroyTaoId).QueryRow(&categoryTao)
	//	var category int = categoryTao.Category
	//	fmt.Println(categoryTao.Category)
	//
	//	var categoryTao2 []CategoryTao
	//	o.Raw("SELECT id FROM category_tao where category=? and id not in(select id from category_tao where id<=?)", category, categroyTaoId).QueryRows(&categoryTao2)
	//	fmt.Println(categoryTao2)
	//
	//	nextcategoryTaoId := categoryTao2[0].Id
	//	fmt.Println(nextcategoryTaoId)
	//	var userLastQuestion UserLastQuestion
	//	userLastQuestion.QueryCategoryTaoId = nextcategoryTaoId
	//	userLastQuestion.UserId = userId
	//	userLastQuestion.Created = time.Now()
	//	userLastQuestion.Updated = time.Now()
	//	o.Insert(&userLastCategoryTao2)
	//
	//} else {
	//	fmt.Println("不相等,有错误或者为答题")
	//}
	fmt.Println("----------------------------------------")
}

// 记录最后的用户的id
func AddUserLastQuestion(userId int, questionId int, categroyTaoId int) {
	fmt.Println("记录最后的用户")
	// fmt.Println("用户编号：" + strconv.Itoa(userId))

	//o := orm.NewOrm()
	//o.Using("default")
	//var userLastQuestion UserLastQuestion
	//
	//err := o.Raw("SELECT id, question_last_id  FROM user_last_question WHERE user_id = ? and question_tao_id = ?", userId, categroyTaoId).QueryRow(&userLastQuestion)
	//if err == nil {
	//	fmt.Println("user nums: ")
	//}
	//var id = userLastQuestion.Id
	//fmt.Println("从数据库拿到的id：" + strconv.Itoa(id))
	//fmt.Println(userLastQuestion.QuestionLastId)
	//if id == 0 {
	//	var userLastQuestion UserLastQuestion
	//	userLastQuestion.QuestionTaoId = categroyTaoId
	//	userLastQuestion.QuestionLastId = questionId
	//	userLastQuestion.UserId = userId
	//	userLastQuestion.Created = time.Now()
	//	userLastQuestion.Updated = time.Now()
	//
	//	o.Insert(&userLastQuestion)
	//} else {
	//	fmt.Println("最后的id:" + strconv.Itoa(questionId))
	//
	//	if questionId != userLastQuestion.QuestionLastId {
	//		userLastQuestion.QuestionLastId = questionId
	//	}
	//
	//	userLastQuestion.Updated = time.Now()
	//
	//	o.Update(&userLastQuestion, "Updated", "QuestionLastId")
	//}
}

// 记录最后用户的套题id
func GetUserLastQuestionId(userId int, categroyId int, categroy2 int) Question {
	fmt.Println("-------------------GetUserLastQuestionId start---------------------")
	var question_id int
	beego.Informational(userId)

	o := orm.NewOrm()
	o.Using("default")
	//
	var userErrorQuestion UserErrorQuestion
	o.Raw("SELECT  question_id   FROM user_error_question WHERE user_id = ? and category_tao_id = ? and categroy = ? order by question_id Desc", userId, categroyId, categroy2).QueryRow(&userErrorQuestion)
	//
	//beego.Informational(userErrorQuestion)
	//var questionLastId = userLastQuestion.QuestionLastId
	//beego.Informational(questionLastId)
	//
	question_id = userErrorQuestion.QuestionId
	var question Question
	o.Raw("SELECT id, num  FROM question WHERE id = ? ", question_id).QueryRow(&question)
	//beego.Informational(question)
	//fmt.Println("---------------------GetUserLastQuestionId end-------------------")

	return question
}

// 答错的题记录
func GetQuestionError(userId int, per int, pagesize int, typeInt int, lastQuestionId int, categroy int) QuestionList {
	fmt.Println("-----------------------------------------")
	fmt.Println("[categroy]:", categroy)
	var questionList QuestionList

	o := orm.NewOrm()
	offset := 0

	// 拿到用户访问的最大记录，数据从最大记录开始

	fmt.Println("[per]:", per)
	if per == 0 || per == 1 {
		offset = 0
	} else if per > 1 {
		offset = (per - 1) * pagesize
	}
	fmt.Println("[offset]:", offset)
	fmt.Println("[pagesize]:", pagesize)

	// 拿到错误题
	var userErrorQuestion []UserErrorQuestion
	o.Raw("SELECT question_id, category_tao_id FROM user_error_question where user_id=? and status=? and categroy = ? ", userId, 1, categroy).QueryRows(&userErrorQuestion)
	if len(userErrorQuestion) == 0 {
		return questionList
	}
	categoryTaoId := userErrorQuestion[0].CategoryTaoId
	fmt.Println(categoryTaoId)
	fmt.Println("[userErrorQuestion]:")
	fmt.Println(userErrorQuestion)
	if len(userErrorQuestion) == 0 {
		return questionList
	}
	errorIds := []int{}
	for _, v1 := range userErrorQuestion {
		errorIds = append(errorIds, v1.QuestionId)
	}
	fmt.Println("[所有的错误的记录]:", errorIds)

	var question []Question
	//num, err := o.Raw("SELECT id,num,title,img ,answer_analysis  FROM question WHERE category_tao=? ORDER BY num limit ?,? ", id, offset, pagesize).QueryRows(&question)
	o.QueryTable("question").Filter("id__in", errorIds).OrderBy("id").Limit(pagesize, offset).All(&question)

	fmt.Println("[question]:")
	fmt.Println(question)
	if len(question) == 0 {
		return questionList
	}

	// 拿到Num字段
	ids := []int{}
	for _, v := range question {
		ids = append(ids, v.Num)
	}
	// s := strings.TrimRight(inData, ",")
	fmt.Println("[当前题的编号]:", ids)

	var answer []Answer
	//
	o.QueryTable("answer").Filter("CategoryTao", categoryTaoId).Filter("num__in", ids).All(&answer)
	fmt.Println("[answer]:")
	fmt.Println(answer)
	// 最终返回
	var reQuestion []ReQuestion
	for _, v := range question {
		num := v.Num
		var newAnswer []Answer
		var optionType int
		for _, vv := range answer {
			if num == vv.Num {
				newAnswerRow := Answer{Num: vv.Num, OptionName: vv.OptionName, OptionAnswer: vv.OptionAnswer, OrderNum: vv.OrderNum}
				newAnswer = append(newAnswer, newAnswerRow)
			}

		}
		for _, vvv := range newAnswer {
			if vvv.OptionAnswer == 1 {
				optionType++
			}
		}
		newStruct := ReQuestion{Id: v.Id, Num: v.Num, Title: v.Title, Img: v.Img, AnswerAnalysis: v.AnswerAnalysis, Answer: newAnswer, OptionType: optionType}
		reQuestion = append(reQuestion, newStruct)

	}
	var tmpPagesize = float64(len(userErrorQuestion)) / float64(pagesize)
	beego.Informational(tmpPagesize)
	questionList = QuestionList{ReQuestion: reQuestion, Total: float64(len(userErrorQuestion)), TotalPage: math.Ceil(tmpPagesize), UserId: userId, CategoryTaoId: categoryTaoId}
	fmt.Println("-----------------------------------------")
	return questionList
}

func GetUserSessionKey(id int) string {

	var user User
	o := orm.NewOrm()
	o.Raw("SELECT sessionkey FROM user where id=? ", id).QueryRow(&user)
	return user.Sessionkey
}

func GetUserList(userType int) []User {
	var user []User
	o := orm.NewOrm()
	o.Raw("SELECT id, mobile_num, user_type, user_name, position, nick_name, created FROM user where user_type=?", userType).QueryRows(&user)
	return user

}

func GetMyOrder(userId int) []UserCategory {
	var userCategory []UserCategory
	o := orm.NewOrm()
	o.Raw("SELECT category_id FROM user_category where user_id=?", userId).QueryRows(&userCategory)
	return userCategory
}

func SubmitMyOrder(userId int, categoryIds string) {
	o := orm.NewOrm()
	o.Using("default")

	arr := []string{}
	if categoryIds != "" {

		arr = strings.Split(categoryIds, ",")
	}
	fmt.Println("userId:" + strconv.Itoa(userId))
	fmt.Println(arr)

	var user User
	user.Id = userId
	user.UserType = 2
	user.Updated = time.Now()
	o.Update(&user, "Updated", "UserType")

	o.Raw("DELETE FROM user_category WHERE user_id= ?", userId).Exec();
	fmt.Println("目前有：" + strconv.Itoa(len(arr)))
	if len(arr) > 0 {

		for _, v := range arr {
			fmt.Println(v)

			var userCategory UserCategory
			userCategory.UserId = userId
			userCategory.Created = time.Now()
			id, _ := strconv.Atoi(v)
			userCategory.CategoryId = id
			o.Insert(&userCategory)

		}
	}

}


func SendSms(mobile, code string) {

	o := orm.NewOrm()
	o.Using("default")

	var tempSms Sms


	o.Raw("SELECT id, code FROM sms WHERE mobile=?  ", mobile).QueryRow(&tempSms)

	beego.Informational(tempSms.Code)
	if tempSms.Code == ""{
		beego.Informational("空")
		var sms Sms
		sms.Mobile = mobile
		sms.Code = code
		sms.Created = time.Now()
		o.Insert(&sms)
	} else{
		beego.Informational("不空")
		tempSms.Mobile = mobile
		tempSms.Code = code
		tempSms.Updated = time.Now()
		o.Update(&tempSms, "Mobile", "Code", "Updated")
	}

	util.SendSms(mobile,code)
}

func SubmitUserInfo(userId int , reCode string, username string, position string, mobileNum string) string{
	var str string
	o := orm.NewOrm()
	o.Using("default")


	var sms Sms
	o.Raw("SELECT id, code FROM sms WHERE mobile=?  ", mobileNum).QueryRow(&sms)

	if sms.Code == reCode {
		var user User
		user.Id = userId
		user.UserName = username
		user.Position = position
		user.MobileNum = mobileNum
		user.Updated = time.Now()
		o.Update(&user, "UserName", "Position", "MobileNum", "Updated")

		o.Raw("DELETE FROM sms WHERE id= ?", sms.Id).Exec();
		str = "succ"
	} else {
		str = "fail"
	}

	return str
}

func CheckUserData(userId int) string {
	var str string
	o := orm.NewOrm()
	o.Using("default")


	var user User
	o.Raw("SELECT mobile_num FROM user WHERE id=?  ", userId).QueryRow(&user)
	if user.MobileNum != ""{
		str = "fail"
	}
	return str
}
