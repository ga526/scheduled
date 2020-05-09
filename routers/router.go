package routers

import (
	"github.com/astaxie/beego/plugins/cors"
	"scheduled/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		AllowOrigins: []string{"*"},
	}))

    beego.Router("/", &controllers.MainController{})
	{
		//导入字典
		beego.Router("/stop", &controllers.Task{}, "*:Stop")
		beego.Router("/start", &controllers.Task{}, "*:Start")
		beego.Router("/todo", &controllers.Task{}, "*:Todo")
	}
}