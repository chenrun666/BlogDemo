package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 账户表
type Account struct {
	Id       int
	Name     string
	Password string
	Token    string
	Userinfo *Userinfo `orm:"reverse(one)"`
}

// 用户信息表
type Userinfo struct {
	Id     int
	Age    int16
	Gender bool
	Phone  string
	Qq     string

	Account *Account   `orm:"rel(one)"`
	Article []*Article `orm:"reverse(many)"`
}

// 文章内容表
type Article struct {
	Id     int
	Title  string
	Brief  string
	Atime  time.Time `orm:"type(datetime);auto_now_add"`
	Browse int       `orm:"default(0)"`
	Img    string    `orm:"null"`
	Text   string    `orm:"type(text)"`

	Userinfo *Userinfo `orm:"rel(fk);null"` // 设置一对多关系
	Category *Category `orm:"rel(fk);null"`

	Comment []*Comment `orm:"reverse(many);null"`
}

// 文章分类
type Category struct {
	Id    int
	Title string

	Article []*Article `orm:"reverse(many)"`
}

// comment评论表
type Comment struct {
	Id      int
	Content string

	Article *Article `orm:"rel(fk)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/blogDemo?charset=utf8")
	orm.RegisterModel(new(Account), new(Userinfo), new(Article), new(Category), new(Comment))
	orm.RunSyncdb("default", false, true)
}
