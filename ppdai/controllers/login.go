package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/xujianhuaap/gostudy/ppdai/api"
	"github.com/astaxie/beego/session"
	"database/sql"
)
var globalSessions *session.Manager
type LoginController struct {
	beego.Controller
}

func ( c *LoginController) Get()  {

	//sess,_:=globalSessions.SessionStart(c.Ctx.ResponseWriter,c.Ctx.Request)
	//defer sess.SessionRelease(c.Ctx.ResponseWriter)
	//username:=sess.Get("username")
	//if(username==nil){
	//
	//}else {
	//	sess.Set("username",c.Ctx.Request.Form["username"])
	//}
	var url=c.Ctx.Request.URL.Path
	if strings.EqualFold(url,api.LOGIN){
		c.TplName="login.html"
	}else if strings.EqualFold(url,api.REGISTER){
		c.TplName="register.html"
	}
}

func (c *LoginController)Post()  {
	db, err:= sql.Open("mysql", "root:123456@/ppdai")
	if err==nil{
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
			stmt,err:=db.Prepare("INSERT user SET name=?,password=?,age=?,gender=?")
			if(err!=nil){
				c.Ctx.WriteString(err.Error())
			}else {
				_,err=stmt.Exec(username,userpassword,26,1)

				if(err!=nil){
					c.Ctx.WriteString("注册失败")
				}else {
					c.Ctx.WriteString("注册成功")
				}

			}
		}
	}else {
		c.Ctx.WriteString("22222")
	}
}


//func init() {
//
//		//conf := new(session.ManagerConfig)
//		//
//		//conf.CookieName = beego.BConfig.WebConfig.Session.SessionName
//		//conf.EnableSetCookie = beego.BConfig.WebConfig.Session.SessionAutoSetCookie
//		//conf.Gclifetime = beego.BConfig.WebConfig.Session.SessionGCMaxLifetime
//		//conf.Secure = beego.BConfig.Listen.EnableHTTPS
//		//conf.CookieLifeTime = beego.BConfig.WebConfig.Session.SessionCookieLifeTime
//		//conf.ProviderConfig = filepath.ToSlash(beego.BConfig.WebConfig.Session.SessionProviderConfig)
//		////conf.Domain = beego.WebConfig.Session.SessionDomain
//		//conf.EnableSidInHttpHeader = beego.BConfig.WebConfig.Session.EnableSidInHttpHeader
//		//conf.SessionNameInHttpHeader = beego.BConfig.WebConfig.Session.SessionNameInHttpHeader
//		//conf.EnableSidInUrlQuery = beego.BConfig.WebConfig.Session.EnableSidInUrlQuery
//		//globalSessions, _ = session.NewManager("mysql", conf)
//
//	}