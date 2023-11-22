package main

import (
	"crudDemo/routers"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	routers.RoutersFunction()
	beego.Run()

}
