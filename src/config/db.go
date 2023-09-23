package config

import (
	"gin_golang/src/models"

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
	DB.AutoMigrate(&models.Article{})
}
