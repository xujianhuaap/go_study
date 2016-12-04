package main

import (
	"github.com/astaxie/beego"

	_"github.com/xujianhuaap/go_studty/ppdai/routers"
)
func main() {
	beego.Run("127.0.0.1:8080")
}

func init() {
	beego.BConfig.WebConfig.DirectoryIndex=true
	beego.SetStaticPath("/apk","apk")
}