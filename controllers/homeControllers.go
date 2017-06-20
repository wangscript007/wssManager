package controllers

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"wssManager/backendCtrl"

	"github.com/astaxie/beego"
)

const tokenName = "usrData"

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	session := this.GetSession(tokenName)
	if nil == session {
		//		this.TplName = "login.html"
		this.TplName = "homecontroller/get.html"
		//		this.Data["Errcode"] = "o123"
	} else {
		this.TplName = "homecontroller/get.html"
	}
}

func (this *HomeController) Post() {
	log.Println("login to service")
	usrName := this.GetString("usrName")
	passWord := this.GetString("pwd")
	addr := this.GetString("addr")
	err := this.login(addr, usrName, passWord)
	if err != nil {
		//登录失败，返回登录界面，并显示错误码
		this.TplName = "homecontroller/get.html"
		this.Data["Errcode"] = err.Error() + "  please login again."
		//this.Redirect("/home", 302)
		return
	} else {
		//重定向
		this.Redirect("/livingList", 302)
	}
}

func (this *HomeController) login(addr, usrName, pwd string) (err error) {
	reqAddr := "http://" + addr + "/admin/login"

	data := make(url.Values)

	data["username"] = []string{usrName}
	data["password"] = []string{pwd}
	resp, err := http.PostForm(reqAddr, data)
	defer func() {
		if err != nil {
			this.DelSession(tokenName)
		}
	}()
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

	beego.Debug(string(body))
	respData, err := backendCtrl.ParseRespData(body)
	if err != nil {
		beego.Debug(err.Error())
		return
	}
	if respData.Code != 200 {
		err = errors.New("login failed:" + respData.Msg)
		return
	}
	usrData := respData.Data.UserData
	this.SetSession(tokenName, usrData)
	return
}
