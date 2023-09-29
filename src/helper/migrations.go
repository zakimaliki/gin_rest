package helper

import (
	"gin_golang/src/config"
	"gin_golang/src/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Article{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	config.DB.AutoMigrate(&models.User{}).Related(&models.Article{})
}
