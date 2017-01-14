package routers

import (
	"github.com/astaxie/beego"
	"github.com/xujianhuaap/gostudy/ppdai/controllers"
	"github.com/xujianhuaap/gostudy/ppdai/api"
)

func init() {
    	beego.Router("/", &controllers.MainController{})
	beego.Router(api.REGISTER,&controllers.LoginController{},"get:Get")
	beego.Router(api.REGISTER_SUBNIT,&controllers.LoginController{},"post:Post")
	beego.Router(api.LOGIN,&controllers.LoginController{},"get:Get")
	beego.Router(api.LOGIN_AUTH,&controllers.LoginController{},"post:Post")
	beego.Router(api.INDEX,&controllers.BusinessController{})

}
