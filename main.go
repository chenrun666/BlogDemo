package main

import (
	_ "blogDemo/routers"
	"github.com/astaxie/beego"
	_ "blogDemo/models"
)

func main() {
	beego.SetStaticPath("/static/css/login/", "/static/")
	beego.SetStaticPath("/static/css/font-awesome-4.7.0/css/", "/static/")
	beego.SetStaticPath("/static/css/article/", "/static/")

	beego.Run()
}

