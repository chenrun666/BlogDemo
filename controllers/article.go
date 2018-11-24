package controllers

import "github.com/astaxie/beego"

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) ShowArticle() {
	c.TplName = "article.html"
}
