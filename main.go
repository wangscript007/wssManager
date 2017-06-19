package main

import (
	_ "wssManager/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/css", "static/css")
	beego.Run()
}
