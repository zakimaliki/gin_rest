package models

import (
	"gin_golang/src/config"

	"github.com/danilopolani/gocialite/structs"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	FullName string
	Email    string
	SocialId string
	Provider string
	Avatar   string
	Role     bool `gorm:"default:0"`
}

func GetOrRegisterUser(provider string, user *structs.User) User {
	var userData User

	config.DB.Where("provider = ? AND social_id = ?", provider, user.ID).First(&userData)

	if userData.ID == 0 {
		newUser := User{
			FullName: user.FullName,
			Email:    user.Email,
			SocialId: user.ID,
			Avatar:   user.Avatar,
		}
		config.DB.Create(&newUser)
		return newUser
	} else {
		return userData
	}
}
