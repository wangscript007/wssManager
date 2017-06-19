package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

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
		addr := this.GetString("addr")
		this.login(addr, usrName, passWord)
	} else {
		log.Println("find session")
	}
}

func (this *HomeController) login(addr, usrName, pwd string) (err error) {
	reqAddr := "http://" + addr + "/admin/login"

	data := make(url.Values)

	data["username"] = []string{usrName}
	data["password"] = []string{pwd}
	resp, err := http.PostForm(reqAddr, data)
	if err != nil {
		beego.Debug(err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err.Error())
		return
	}
	beego.Debug(pwd)
	beego.Debug(string(body))
	return
}
