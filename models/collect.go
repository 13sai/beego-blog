package models

import (
	"github.com/astaxie/beego/orm"
	// "fmt"
)

type Collect struct {
    Id int `orm:"pk"`
    Title string
	Url string
}

func (m *Collect) TableName() string {
    return "collect"
}

func init(){
	orm.RegisterModel(new(Collect))
}

func GetCollectList() (list []orm.Params) {
	o := orm.NewOrm()
	collect := new(Collect)
	qs := o.QueryTable(collect)
	qs.OrderBy("-id").Values(&list)
	
	return list
}
