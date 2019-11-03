package controllers

import (
	"github.com/astaxie/beego"
    "fmt"
    "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"beego-blog/models"
	"beego-blog/services"
	"strconv"
	"strings"
	"os"
	"github.com/mitchellh/mapstructure"
)

func init() {
	defer func() {
        if info := recover(); info != nil {
			fmt.Println("错误：", info)
			os.Exit(100)
        }
    }()
	var sqlConn strings.Builder
	sqlConn.WriteString(beego.AppConfig.String("mysqluser"))
	sqlConn.WriteString(":")
	sqlConn.WriteString(beego.AppConfig.String("mysqlpw"))
	sqlConn.WriteString("@tcp(")
	sqlConn.WriteString(beego.AppConfig.String("mysqlhost"))
	sqlConn.WriteString(":")
	sqlConn.WriteString(beego.AppConfig.String("mysqlport"))
	sqlConn.WriteString(")/")
	sqlConn.WriteString(beego.AppConfig.String("mysqldb"))
	sqlConn.WriteString("?charset=utf8")
	sqlConnStr := sqlConn.String()

    orm.RegisterDriver("mysql", orm.DRMySQL)
	ret	:= orm.RegisterDataBase("default", "mysql", sqlConnStr)
	if ret != nil {
		panic(ret)
	}
	// 最大连接数
	orm.SetMaxIdleConns("default",100)
	orm.SetMaxOpenConns("default",128)
	// orm.RegisterDataBase("default", "mysql", "root:null@tcp(192.168.1.1:3306)/blog?charset=utf8")
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	key := "article_list"
	list := services.GetCache(key)
	if list == nil {
		list = models.GetArticleList()
		services.SetCache(key, list)
	}
	c.Data["recommends"] = GetRecommend()
	c.Data["articles"] = list
	c.Layout = "home.html"
    c.LayoutSections = make(map[string]string)
    c.LayoutSections["Side"] = "side.html"
    c.TplName = "article.html"
}

func GetRecommend() (ret interface{}) {
	key := "recommend"

	ret = services.GetCache(key)
	if ret == nil {
		ret = models.GetRecommend()
		services.SetCache(key, ret)
	}

	return ret
}

func getArticleCache(id int) (article models.Article) {
	key := fmt.Sprintf("article/%d", id)

	cache := services.GetCache(key)
	if cache == nil {
		article = models.GetArticle(id)
		services.SetCache(key, article)
	} else {
		err := mapstructure.Decode(cache, &article)
		if err != nil {
			fmt.Println(err)
		}
	}

	return article
}

func (this *MainController) Essay() {
	id := this.Ctx.Input.Param(":id")
	iid, _ := strconv.Atoi(id)
	article := getArticleCache(iid)

	this.Data["info"] = article
	fmt.Printf("%T", article)
	this.Data["Title"] = article.Title
	this.Data["recommends"] = GetRecommend()
	this.TplName = "essay.html"
	this.Layout = "home.html"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Side"] = "side.html"
}

func (this *MainController) About() {
	article := getArticleCache(112)

	this.Data["info"] = article
	// this.Data["Title"] = article.Title
	this.Data["recommends"] = GetRecommend()
	this.TplName = "page.html"
	this.Layout = "home.html"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Side"] = "side.html"
}

func (this *MainController) Tag() {
	article := getArticleCache(129)

	this.Data["info"] = article
	// this.Data["Title"] = article.Title
	this.Data["recommends"] = GetRecommend()
	this.TplName = "page.html"
	this.Layout = "home.html"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Side"] = "side.html"
}

func (this *MainController) Collect() {
	key := "collect"
	list := services.GetCache(key)
	if list == nil {
		list = models.GetCollectList()
		services.SetCache(key, list)
	}
	this.Data["list"] = list
	this.Data["recommends"] = GetRecommend()
	this.TplName = "collect.html"
	this.Layout = "home.html"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Side"] = "side.html"
}
