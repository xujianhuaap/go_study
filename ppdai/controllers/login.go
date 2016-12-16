package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/xujianhuaap/gostudy/ppdai/api"
	"encoding/base64"
)

type LoginController struct {
	beego.Controller
}

func ( c *LoginController) Get()  {
	secret:=base64.StdEncoding.EncodeToString("ppdai123")
	cookie,_:=c.GetSecureCookie(secret,"login")
	if(cookie){

	}else {
		c.SetSession()
	}
	var url=c.Ctx.Request.URL.Path
	if strings.EqualFold(url,api.LOGIN){
		c.TplName="login.html"
	}else if strings.EqualFold(url,api.REGISTER){
		c.TplName="register.html"
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
			c.Ctx.Redirect(200,api.LOGIN)
		}
	}else if strings.EqualFold(url,api.REGISTER_SUBNIT){
		username:=c.GetString("user_name","")
		userpassword:=c.GetString("user_password","")

		if(len(username)>0&& len(userpassword)>8){
			c.Ctx.WriteString("注册成功")
		}else {
			c.Ctx.WriteString("注册失败")
		}
	}


}

