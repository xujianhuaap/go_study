package main

import (
	"github.com/astaxie/beego"
	_"github.com/xujianhuaap/gostudy/ppdai/routers"
)

func main() {
	beego.Run()


}

type SessionData struct {
	Token string
}

func init() {

	beego.BConfig.WebConfig.EnableXSRF=false
	beego.BConfig.WebConfig.XSRFKey="ppdai123456"
	beego.BConfig.WebConfig.XSRFExpire=3600
	beego.BConfig.WebConfig.DirectoryIndex=true
	beego.BConfig.WebConfig.ViewsPath="views"

	beego.BConfig.WebConfig.Session.SessionOn=true//
	beego.BConfig.WebConfig.Session.SessionProvider="mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig="root:123456@/ppdai"
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime=180
	beego.BConfig.WebConfig.Session.SessionDomain="192.168.51.163"
	beego.BConfig.WebConfig.Session.SessionName = "ppdai"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true
	beego.BConfig.WebConfig.Session.EnableSidInHttpHeader=true
	beego.BConfig.WebConfig.Session.EnableSidInUrlQuery=true
	beego.BConfig.WebConfig.Session.SessionNameInHttpHeader="Ppdai"

	beego.SetStaticPath("/apk","apk")

}