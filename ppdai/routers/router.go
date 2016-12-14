package routers

import (
	"github.com/astaxie/beego"
	"github.com/xujianhuaap/gostudy/ppdai/controllers"
)

func init() {
    	beego.Router("/", &controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{},"get:Login")
}
