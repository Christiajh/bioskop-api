package main

import (
	"bioskop-api/config"
	"bioskop-api/routers"
)

func main() {
	config.ConnectDB()
	r := routers.SetupRouter()
	r.Run(":8080")
}
