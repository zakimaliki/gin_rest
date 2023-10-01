package helper

import (
	"gin_golang/src/config"
	"gin_golang/src/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Article{})
}
