package main

import (
	"gin_golang/src/config"
	"gin_golang/src/models"
	"gin_golang/src/routes"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	config.DB.AutoMigrate(&models.Article{})
	defer config.DB.Close()
	routes.Routing()
}
