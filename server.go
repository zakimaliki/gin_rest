package main

import (
	"gin_golang/src/config"
	"gin_golang/src/models"
	"gin_golang/src/routes"
)

func main() {
	config.InitDB()
	config.DB.AutoMigrate(&models.Article{})
	defer config.DB.Close()
	routes.Routing()
}
