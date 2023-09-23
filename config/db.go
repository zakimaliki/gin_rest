package config

import (
	"gin_golang/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// defer DB.Close()

	// Migrate the schema
	DB.AutoMigrate(&models.Article{})
}
