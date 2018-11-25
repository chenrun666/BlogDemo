package main

import (
	"blogDemo/controllers"
	_ "blogDemo/routers"
	"github.com/astaxie/beego"
	_ "blogDemo/models"
)

func main() {
	// 静态文件
	beego.SetStaticPath("/static/css/login/", "/static/")
	beego.SetStaticPath("/static/css/font-awesome-4.7.0/css/", "/static/")
	beego.SetStaticPath("/static/css/article/", "/static/")
	beego.SetStaticPath("/static/js/bootstrap-paginator-master/build/", "/static/")

	// 启用session
	beego.BConfig.WebConfig.Session.SessionOn = true

	// 模板函数的映射
	beego.AddFuncMap("ShowPrevious", controllers.HandlePrevious)
	beego.AddFuncMap("ShowNext", controllers.HandleNext)

	beego.Run()
}

