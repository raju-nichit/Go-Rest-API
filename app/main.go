package main

import (
	"Go-Rest-API/server"
	"go-rest-api/config"
)

func main() {
	config.DBConfig()
	server.Init()
}
