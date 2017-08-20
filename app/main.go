package main

import (
	"github.com/raju-nichit/Go-Rest-API/config"
	"github.com/raju-nichit/Go-Rest-API/server"
)

func main() {
	config.DBConfig()
	server.Init()
}
