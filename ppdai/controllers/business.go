package controllers

import (
	"github.com/astaxie/beego"
)

type BusinessController struct {
	beego.Controller
}


func ( c *BusinessController) Get()  {
	responseWriter:=c.Ctx.ResponseWriter.ResponseWriter
	request:=c.Ctx.Request
	sess,err:=beego.GlobalSessions.SessionStart(responseWriter,request)
	beego.BeeLogger.Debug("session sid :%v",sess.SessionID())
	if(err==nil&&sess!=nil){
		data:=sess.Get("user")
		beego.BeeLogger.Debug("session_data :%+v",data)
		if(data!=nil){
			c.TplName="family.html"
			if err!=nil{
				panic(err)
			}

		}else {
			if(err!=nil){
				beego.BeeLogger.Debug("err :%v",err.Error())
			}
			c.Ctx.WriteString("请重新登录")
		}


	}else {
		beego.BeeLogger.Debug("err :%v",err.Error())
	}

}


func init() {

}
