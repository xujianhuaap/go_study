package main

import (
	"github.com/astaxie/beego"

	_"github.com/xujianhuaap/gostudy/ppdai/routers"
)
func main() {
	beego.Run("192.168.23.92:8080")
}

func init() {
	beego.BConfig.WebConfig.EnableXSRF=false
	beego.BConfig.WebConfig.XSRFKey="ppdai123456"
	beego.BConfig.WebConfig.XSRFExpire=3600
	beego.BConfig.WebConfig.DirectoryIndex=true
	beego.BConfig.WebConfig.ViewsPath="views"
	beego.BConfig.WebConfig.Session.SessionOn=true
	beego.BConfig.WebConfig.Session.SessionProvider="mysql"
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime=0//用户关闭浏览器的时候就可以关闭cookie了
	beego.SetStaticPath("/apk","apk")

}