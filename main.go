package main

import (
	_ "beego-blog/routers"
	"github.com/astaxie/beego"
)

func init() {}

func main() {
	beego.SetStaticPath("/static","static")
	beego.Run()
}

