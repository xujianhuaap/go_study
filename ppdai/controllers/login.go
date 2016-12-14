package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/xujianhuaap/gostudy/ppdai/api"
)

type LoginController struct {
	beego.Controller
}

func ( c *LoginController) Get()  {
	var url=c.Ctx.Request.URL.Path
	if strings.EqualFold(url,api.LOGIN){
		c.TplName="login.html"
	}
}

func (c *LoginController)Post()  {
	var url=c.Ctx.Request.URL.Path
	if strings.EqualFold(url,api.LOGIN_AUTH){
		username:=c.GetString("user_name","")
		userpassword:=c.GetString("user_password","")
		if(strings.EqualFold(username,"xu")&& strings.EqualFold(userpassword,"123456")){
			c.Ctx.WriteString("登录成功")
		}else {
			c.Ctx.WriteString("登录失败")
		}
	}


}

