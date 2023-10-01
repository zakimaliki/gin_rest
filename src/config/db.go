package config

import (
	// "os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	// url := os.Getenv("URL")
	var err error
	DB, err = gorm.Open("postgres", "host=147.139.210.135 port=5432 user=zaki dbname=zaki01 password=zaki123")
	if err != nil {
		panic("failed to connect database")
	}
}
