package main

import (
	_"github.com/xujianhuaap/GoStudy/ppdai/routers"
	"github.com/astaxie/beego"
	"github.com/beego/bee/testdata/router"
)
func main() {
	beego.Router("/",&router.Router{})
	beego.BConfig.WebConfig.DirectoryIndex=true
	beego.SetStaticPath("/apk","apk")
	beego.Run()
}

func init() {

}