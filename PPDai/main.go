package main

import (
	"github.com/astaxie/beego"
	_"github.com/xujianhuaap/gostudy/ppdai/routers"

)
func main() {
	beego.Run()
}

func init() {
	beego.BConfig.WebConfig.DirectoryIndex=true
	beego.SetStaticPath("/apk","apk")
}