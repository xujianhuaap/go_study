package main

import (
	"github.com/astaxie/beego"

	_"github.com/xujianhuaap/gostudy/ppdai/routers"
	 _ "github.com/go-sql-driver/mysql"

	"database/sql"
)

var db *sql.DB
func main() {
	beego.Run()


}

func init() {

	beego.BConfig.WebConfig.EnableXSRF=false
	beego.BConfig.WebConfig.XSRFKey="ppdai123456"
	beego.BConfig.WebConfig.XSRFExpire=3600
	beego.BConfig.WebConfig.DirectoryIndex=true
	beego.BConfig.WebConfig.ViewsPath="views"

	//beego.AppConfig.String("mysqluser")="xujianhua"
	//beego.AppConfig.String("mysqlpass")="123456"
	//beego.AppConfig.String("mysqlurl")="127.0.0.1"
	//beego.AppConfig.String("mysqldb")="ppdai"

	//beego.BConfig.WebConfig.Session.SessionOn=true
	//beego.BConfig.WebConfig.Session.SessionProvider="mysql"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig="xujianhua:123456@localhost/ppdai?charset=utf8"
	//beego.BConfig.WebConfig.Session.SessionCookieLifeTime=0//用户关闭浏览器的时候就可以关闭cookie了
	//beego.BConfig.WebConfig.Session.SessionDomain="xjh"

	beego.SetStaticPath("/apk","apk")

}