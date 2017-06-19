package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	session := this.GetSession("usr")
	if nil == session {
		//		this.TplName = "login.html"
		this.TplName = "homecontroller/get.html"
	} else {
		log.Println("find session")
	}
}

func (this *HomeController) Post() {
	session := this.GetSession("usr")
	if nil == session {
		log.Println("login to service")
		usrName := this.GetString("usrName")
		passWord := this.GetString("pwd")
		log.Println(usrName)
		log.Println(passWord)
	} else {
		log.Println("find session")
	}
}
