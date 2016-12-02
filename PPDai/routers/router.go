package routers

import (
	"github.com/xujianhuaap/gostudy/ppdai/controllers"
	"github.com/astaxie/beego"
)

func init() {
    	beego.Router("/", &controllers.MainController{})
	beego.Router("/java",&controllers.FunnyController{},"post:Post")
}
