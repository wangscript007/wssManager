package routers

import (
	"wssManager/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.HomeController{})
}
