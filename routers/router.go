package routers

import (
	"blogDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 主页面
	beego.Router("/", &controllers.MainController{}, "get:Index")
	// 登陆注册
	beego.Router("/login", &controllers.AccountController{}, "get:Login;post:LoginSubmit")
	beego.Router("/register", &controllers.AccountController{}, "get:Register;post:RegisterSubmit")

	// 文章跳转路径
	beego.Router("/article", &controllers.ArticleController{}, "get:ShowArticle")
}
