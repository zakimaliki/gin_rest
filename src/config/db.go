package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	url := os.Getenv("URL")
	var err error
	DB, err = gorm.Open("mysql", url)
	if err != nil {
		panic("failed to connect database")
	}
}
