package main

import (
	"github.com/astaxie/beego"

)
func main() {
	beego.Run()
}

func init() {
	beego.BConfig.WebConfig.DirectoryIndex=true
	beego.SetStaticPath("/apk","apk")
}