package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Info() {

	c.Data["Website"] = "xujianhuaap.io"
	c.Data["Email"] = "xujianhuaap@gmail.com"
	c.TplName = "family.html"
}

 
