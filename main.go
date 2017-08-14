package main

import (
	"go-rest-api/config"
	"go-rest-api/webservices"
)

func main() {
	config.DBConfig()
	webservices.RounterConfig()
}
