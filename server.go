package main

import (
	"gin_golang/src/config"
	"gin_golang/src/routes"
)

func main() {
	config.InitDB()
	defer config.DB.Close()
	routes.Routing()
}
