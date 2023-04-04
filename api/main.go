package main

import (
	"api/libs"
	_ "api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//设置日志
	libs.InitLogger()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
