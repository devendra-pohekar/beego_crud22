package main

import (
	"crudDemo/routers"
	"log"

	"github.com/astaxie/beego"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	routers.RoutersFunction()
	beego.Run()

}
