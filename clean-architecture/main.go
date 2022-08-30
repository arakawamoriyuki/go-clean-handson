package main

import (
	"main/infrastructure/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
