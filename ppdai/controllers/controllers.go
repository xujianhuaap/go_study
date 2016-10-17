package controllers

import "github.com/astaxie/beego"

type FunnyController struct {
	beego.Controller
}

func (c FunnyController)Get()  {
	c.Data["email"]="xujianhuaap@gmail.com"
	c.TplName=""
	c.Ctx.WriteString("JAVA")
}