package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/xujianhuaap/gostudy/ppdai/api"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/mysql"
	"database/sql"
	"path/filepath"
)
type LoginController struct {
	beego.Controller
}

var db *sql.DB
var globalSessions *session.Manager

func ( c *LoginController) Get()  {
	responseWriter:=c.Ctx.ResponseWriter.ResponseWriter
	request:=c.Ctx.Request
	sess,err:=globalSessions.SessionStart(responseWriter,request)
	beego.BeeLogger.Debug("session sid :%v",sess.SessionID())
	if(err==nil&&sess!=nil){
		defer sess.SessionRelease(c.Ctx.ResponseWriter)
		data:=sess.Get("user")
		beego.BeeLogger.Debug("session_data :%+v",data)
		if(data!=nil){
			c.TplName="index.html"
			if err!=nil{
				panic(err)
			}

		}else {
			err=sess.Set("user","123456")
			if(err!=nil){
				beego.BeeLogger.Debug("err :%v",err.Error())
			}
			var url=c.Ctx.Request.URL.Path
			if strings.EqualFold(url,api.LOGIN){
				c.TplName="login.html"
			}else if strings.EqualFold(url,api.REGISTER){
				c.TplName="register.html"
			}
		}


	}else {
		beego.BeeLogger.Debug("err :%v",err.Error())
	}

}

func (c *LoginController)Post()  {
	var err error
	db, err = sql.Open("mysql", "root:123456@/ppdai")
	defer db.Close()
	if err==nil{
		var url=c.Ctx.Request.URL.Path
		if strings.EqualFold(url,api.LOGIN_AUTH){
			username:=c.GetString("user_name","")
			userpassword:=c.GetString("user_password","")

			isRegister:=queryUserIsRegister(username,userpassword,db)

			if(isRegister){
				c.TplName="index.html"
			}else {
				c.Ctx.Redirect(301,api.LOGIN)
			}
		}else if strings.EqualFold(url,api.REGISTER_SUBNIT){
			username:=c.GetString("user_name","")
			userpassword:=c.GetString("user_password","")
			isRegister:=queryUserIsRegister(username,userpassword,db)
			if(!isRegister){
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
			}else {
				c.Ctx.Redirect(301,api.LOGIN)
			}

		}
	}else {
		c.Ctx.WriteString("22222")
	}
}
func queryUserIsRegister(user_name string ,user_password string,db*sql.DB) bool{

	stmt,err:=db.Prepare("select id from user where name=? and password=?")
	if(err!=nil){
		beego.BeeLogger.Debug("stmt error %v",err.Error())
		return false
	}
	rows,err:=stmt.Query(user_name,user_password)
	defer rows.Close()
	if(err!=nil){
		beego.BeeLogger.Debug("stmt error %v",err.Error())
		return false
	}
	if(rows.Next()){
		return true
	}

	return false
}

func init() {
		var err error

		beego.BConfig.WebConfig.Session.SessionOn=true
		beego.BConfig.WebConfig.Session.SessionProvider="mysql"
		beego.BConfig.WebConfig.Session.SessionProviderConfig="root:123456@/ppdai"
		beego.BConfig.WebConfig.Session.SessionCookieLifeTime=180
		beego.BConfig.WebConfig.Session.SessionDomain="http://192.168.31.31:8080"
		beego.BConfig.WebConfig.Session.SessionName = "ppdai"
		beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
		beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true
		beego.BConfig.WebConfig.Session.EnableSidInHttpHeader=true
		beego.BConfig.WebConfig.Session.SessionNameInHttpHeader="Ppdai"

		//gob.RegisterName("session_data",SessionData{})
		conf := new(session.ManagerConfig)

		conf.CookieName = beego.BConfig.WebConfig.Session.SessionName
		conf.EnableSetCookie = beego.BConfig.WebConfig.Session.SessionAutoSetCookie
		conf.Gclifetime = beego.BConfig.WebConfig.Session.SessionGCMaxLifetime
		conf.Maxlifetime=300
		conf.SessionIDLength=32
		conf.Secure =true
		conf.CookieLifeTime = beego.BConfig.WebConfig.Session.SessionCookieLifeTime
		conf.ProviderConfig = filepath.ToSlash(beego.BConfig.WebConfig.Session.SessionProviderConfig)
		conf.Domain = "192.168.31.31"
		conf.EnableSidInHttpHeader = beego.BConfig.WebConfig.Session.EnableSidInHttpHeader
		conf.SessionNameInHttpHeader = beego.BConfig.WebConfig.Session.SessionNameInHttpHeader
		conf.EnableSidInUrlQuery = true
		beego.BeeLogger.Debug("session manager conf: %+v",conf)
		globalSessions, err= session.NewManager("mysql", conf)
		if(err!=nil){
			beego.BeeLogger.Debug("login init: %v",err.Error())
		}else {
			go globalSessions.GC()
		}

	}