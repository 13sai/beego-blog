package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Article struct {
    Id int `orm:"pk"`
    Title string
	Abstract string
	ClickRate int
	IsTop int
	Content string
	Date string
}

type ArticleSum struct {
	Num int
    Month string
}

type Ret struct {
    ArticleSum
    List []orm.Params
}

func (m *Article) TableName() string {
    return "article"
}

func init(){
	orm.RegisterModel(new(Article))
}

func GetArticleList() (list []Ret) {
	o := orm.NewOrm()
	article := new(Article)
	qs := o.QueryTable(article)

	var result []orm.Params
	qs.OrderBy("-id").Values(&result)

	var ret []ArticleSum
	sql := "SELECT count(*) as num,left(`date`, 7) as month FROM %s GROUP BY month ORDER BY month desc"
	sql = fmt.Sprintf(sql, article.TableName())
	o.Raw(sql).QueryRows(&ret)

	var rs []Ret
	var sum = 0
	var row = Ret{}

	for _,monthSum := range ret {
		row.Num = monthSum.Num
		row.Month = monthSum.Month
		row.List = result[sum:sum+row.Num]
		// fmt.Println(result[sum:sum+row.Num])
		rs = append(rs, row)
		sum += row.Num
	}
	return rs
}

func GetRecommend() (result []orm.Params) {
	o := orm.NewOrm()
	article := new(Article)
	qs := o.QueryTable(article)

	// var result []orm.Params
	qs.Filter("is_top", 1).Limit(10).OrderBy("-id").Values(&result)
	return result
}

func GetArticle(id int) (result Article) {
	o := orm.NewOrm()
	result = Article{Id: id}

	err := o.Read(&result)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	}
	
	return result
}