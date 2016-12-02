package controllers

import (
	"github.com/astaxie/beego"
)

type FunnyController struct {
	beego.Controller
}

func (c *FunnyController)Get()  {
	header:=c.Ctx.Request.Header;

	body:=c.Ctx.Request.PostForm.Get("name");
	beego.BeeLogger.Debug("Request header %v,Request body %v",header,body)
	c.Data["Email"]="xujianhuaap@gmail.com"
	c.Data["Temp"]="Temp"
	c.Data["Name"]="name"
	c.TplName="index.html"


}
func (c *FunnyController)Post()  {
	header:=c.Ctx.Request.Header;

	body:=c.GetString("name","undefined");
	beego.BeeLogger.Debug("Request header %v,Request body %v",header,body)
	c.Data["Email"]="xujianhuaap@gmail.com"
	c.Data["Temp"]="Temp"
	c.Data["Name"]="name"
	c.TplName="index.html"


}