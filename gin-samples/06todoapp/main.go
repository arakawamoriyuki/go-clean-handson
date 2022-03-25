package main

import (
	"main/models"
	"main/pkg/setting"
	"main/routers"
)

func init() {
	setting.Setup("conf/development.ini")
	models.Setup()
}

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
