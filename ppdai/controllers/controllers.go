package controllers

import "github.com/astaxie/beego"

type FunnyController struct {
	beego.Controller
}

func (c *FunnyController)Get()  {
	c.Data["Email"]="xujianhuaap@gmail.com"
	c.TplName="index.tpl"
}