package main

import (
	_"PPDai/routers"
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