package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/xujianhuaap/gostudy/ppdai/api"
	_ "github.com/astaxie/beego/session/mysql"
	"database/sql"
)
type LoginController struct {
	beego.Controller
}
func ( c *LoginController) Get()  {


	var url=c.Ctx.Request.URL.Path
	if strings.EqualFold(url,api.LOGIN){
		c.TplName="login.html"
	}else if strings.EqualFold(url,api.REGISTER){
		c.TplName="register.html"
	}


}

func (c *LoginController)Post()  {

	responseWriter:=c.Ctx.ResponseWriter.ResponseWriter
	request:=c.Ctx.Request

	db, err := sql.Open("mysql", "root:123456@/ppdai")
	defer db.Close()
	if err==nil{
		var url=c.Ctx.Request.URL.Path
		username:=c.GetString("user_name","")
		userpassword:=c.GetString("user_password","")
		isRegister:=queryUserIsRegister(username,userpassword,db)
		beego.BeeLogger.Debug("user_name %v,user_password %v,isRegister %v",username,userpassword,isRegister)
		if strings.EqualFold(url,api.REGISTER_SUBNIT){
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
				c.Ctx.Redirect(301,api.INDEX)
			}

			sess,err:=beego.GlobalSessions.SessionStart(responseWriter,request)
			defer sess.SessionRelease(responseWriter)
			beego.BeeLogger.Debug("session sid :%v",sess.SessionID())
			if(err==nil&&sess!=nil){

				err=sess.Set("user",username)
				beego.BeeLogger.Debug("user [%v] isRegistered",username)
				if(err!=nil){
					beego.BeeLogger.Debug("err :%v",err.Error())
				}
			}else {
				beego.BeeLogger.Debug("err :%v",err.Error())
			}

		}

		if strings.EqualFold(url,api.LOGIN_AUTH){
			if(!isRegister){
				c.Ctx.Redirect(301,api.REGISTER)
			}else {

				sess,err:=beego.GlobalSessions.SessionStart(responseWriter,request)
				defer sess.SessionRelease(responseWriter)
				beego.BeeLogger.Debug("session sid :%v",sess.SessionID())
				if(err==nil&&sess!=nil){
					err=sess.Set("user",username)
					beego.BeeLogger.Debug("user [%v] isRegistered",username)
					if(err!=nil){
						beego.BeeLogger.Debug("err :%v",err.Error())
					}
				}else {
					beego.BeeLogger.Debug("err :%v",err.Error())
				}
				c.Ctx.Redirect(301,api.INDEX)
			}

		}

	}else {
		c.Ctx.WriteString("22222")
	}


}
func queryUserIsRegister(user_name string ,user_password string,db*sql.DB) bool{

	row:=db.QueryRow("select id from user where `name` = ?",user_name)
	var id int16
	if(row!=nil){
		row.Scan(&id)
		beego.BeeLogger.Debug("queryIsRegister User Id %v",id)
		if(id>0){
			return true
		}
	}

	return false
}

func init() {

}