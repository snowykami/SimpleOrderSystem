package main

import (
	"main/api"
	"main/server"
)

func main() {

	server.StartServer()
	api.Logger.Println("INFO: " + "main")
}
