package controllers

import (
	"blogDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ManageController struct {
	beego.Controller
}

func (this *ManageController) AddArticle() {
	// 获取orm对象
	o := orm.NewOrm()
	var category []models.Category
	_, err := o.QueryTable("Category").All(&category)
	if err != nil{
		fmt.Println("err")
	}else{
		this.Data["category"] = category
	}
	this.TplName = "manage/addarticle.html"
}

func (this *ManageController) AddArticleSubmit() {
	var Art models.Article
	if err := this.ParseForm(&Art); err != nil {
		fmt.Println("出错误了！！！")
	}else {
		// 获取orm对象
		o := orm.NewOrm()
		article := models.Article{}
		article.Title = Art.Title
		article.Brief = Art.Brief
		article.Text = Art.Text
		article.Category = Art.Category
		// 没有添加外间关系，用户的的关联，分组的关联
		_, err := o.Insert(&article)
		if err != nil{
			fmt.Println(err)
			this.Data["errormsg"] = "添加失败"
		}else{
			this.Data["msg"] = "添加成功"
		}

	}

	this.TplName = "manage/addarticle.html"
}
