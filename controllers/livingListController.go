package controllers

import (
	"wssManager/backendCtrl"

	"github.com/astaxie/beego"
)

type LivingListController struct {
	beego.Controller
}

func (this *LivingListController) Get() {
	session := this.GetSession(tokenName)
	if nil == session {
		//		this.TplName = "login.html"
		this.TplName = "homecontroller/get.html"
		//		this.Data["Errcode"] = "o123"
		this.Redirect("/home", 302)
	} else {
		this.TplName = "livingList/livingList.html"
		usr := session.(backendCtrl.Usr)
		this.Data["UsrName"] = usr.Usrname
	}
}
