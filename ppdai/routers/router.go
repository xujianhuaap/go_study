package routers

import (
	"github.com/astaxie/beego"
	"github.com/xujianhuaap/go_studty/ppdai/controllers"
)

func init() {
    	beego.Router("/", &controllers.MainController{})
}
