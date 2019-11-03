package routers

import (
	"beego-blog/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/collect", &controllers.MainController{}, "get:Collect")
	beego.Router("/tag", &controllers.MainController{}, "get:Tag")
	beego.Router("/about", &controllers.MainController{}, "get:About")
	beego.Router("/essay/:id", &controllers.MainController{}, "get:Essay")
	beego.Get("/test", func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
   })
}
