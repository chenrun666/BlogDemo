package controllers

import (
	"blogDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

type MainController struct {
	beego.Controller
}

type AccountController struct {
	beego.Controller
}

func (c *MainController) Index() {
	// 获取当前请求会话，并返回当前请求会话的对象
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	username := sess.Get("username")
	fmt.Printf("username:%v\n", username)
	if username == nil {
		c.Redirect("/login", 302)
	}
	// 返回当前用户的username
	c.Data["username"] = username

	// 获取orm对象，查找数据库，获取分类信息
	o := orm.NewOrm()
	// 存放数据的
	var art []*models.Article
	// 指定查询的表单
	article := models.Article{}
	// 获取categoryId
	categoryId, _ := c.GetInt("category")
	// 指定查询的表
	qsArticle := o.QueryTable(&article)

	// 获取所有分类
	category := models.Category{}
	qs := o.QueryTable(&category)
	var maps []orm.Params
	_, err := qs.Values(&maps)
	if err == nil {
		// 分类
		c.Data["category"] = maps
	}
	// 循环对应点击的分类
	for _, v := range maps {
		if int64(categoryId) == v["Id"] {
			qsArticle.Filter("category__id", categoryId).RelatedSel().All(&art)
			break
		}
	}
	if categoryId == 0 {
		qsArticle.RelatedSel().All(&art)
	}

	// 获取所有的数据总数
	count := len(art)
	fmt.Printf("count: %v\n", count)
	// 获取当前点击的页码, 在数据中查询就是以这个开始，以currentNum+每页展示个数结束。limit限制
	currentNum, _ := c.GetInt("page")
	paginator := PageUtil(count, currentNum, 1, art)


	c.Data["paginator"] = paginator
	c.Data["atricleItem"] = art
	c.Layout = "layout.html"
	c.TplName = "index.html"
}

func (c *AccountController) Login() {
	c.TplName = "login.html"
}

func (c *AccountController) LoginSubmit() {
	// 获取用户名和密码
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println(username, password)

	// 校验数据是否为空
	if username == "" || password == "" {
		c.Data["msg"] = "用户名密码不能为空"
		c.TplName = "login.html"
		return
	}

	// 获取数据库以及orm进行数据库查询
	o := orm.NewOrm()
	account := models.Account{}
	account.Name = username
	err := o.Read(&account, "Name")
	if err != nil {
		fmt.Println(err)
		c.Data["namemsg"] = "没有该用户，请重新登录"
		c.TplName = "login.html"
		return
	}
	if account.Password != password {
		c.Data["pwdmsg"] = "密码错误， 请重新登陆"
		c.TplName = "login.html"
		return
	}
	// 获取当前请求会话，并返回当前请求会话的对象
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// 根据当前请求对象，设置一个session
	sess.Set("username", username)
	c.Redirect("/", 302)

}

func (c *AccountController) LoginOut() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	sess.Delete("username")
	c.Redirect("/login", 302)
}

func (c *AccountController) Register() {
	c.TplName = "register.html"
}

func (c *AccountController) RegisterSubmit() {
	// 获取提交的数据
	username := c.GetString("username")
	password := c.GetString("password")
	passwordagain := c.GetString("pwdagain")

	// 判断接受的数据不能为空
	if username == "" || password == "" || passwordagain == "" {
		c.Data["msg"] = "数据不能为空"
		c.TplName = "register.html"
		return
	}

	// 校验两次密码一致
	if passwordagain != password {
		c.Data["retainusername"] = username
		c.Data["againmsg"] = "两次密码不一致"
		c.TplName = "register.html"
		return
	}
	// 获取orm对象，查询数据库
	o := orm.NewOrm()
	account := models.Account{}
	qs := o.QueryTable(&account)
	obj := qs.Filter("Name", username)
	result := obj.Exist()
	if result {
		c.Data["namemsg"] = "用户名已存在"
		c.TplName = "register.html"
		return
	} else {
		// 添加数据
		account.Name = username
		account.Password = password
		_, err := o.Insert(&account)
		if err != nil {
			c.Data["failed"] = "注册失败，请重新登陆"
			c.TplName = "register.html"
			return
		} else {
			c.Data["success"] = "注册成功，请登陆！"
			c.TplName = "register.html"
			return
		}
	}

	//
	c.TplName = "register.html"

}
