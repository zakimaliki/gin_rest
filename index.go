package main

import (
	"gin_golang/src/config"
	"gin_golang/src/helper"
	"gin_golang/src/routes"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migrate()
	defer config.DB.Close()
	routes.Router()
}
