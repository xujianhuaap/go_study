package routers

import (
	"github.com/astaxie/beego"
	"github.com/xujianhuaap/gostudy/ppdai/controllers"
	"github.com/xujianhuaap/gostudy/ppdai/api"
)

func init() {
    	beego.Router("/", &controllers.MainController{})
	beego.Router(api.REGISTER,&controllers.LoginController{})
	beego.Router(api.REGISTER_SUBNIT,&controllers.LoginController{})
	beego.Router(api.LOGIN,&controllers.LoginController{})
	beego.Router(api.LOGIN_AUTH,&controllers.LoginController{})

}
