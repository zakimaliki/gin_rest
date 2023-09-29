package helper

import (
	"fmt"
	"gin_golang/src/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user *models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"FullName":  user.FullName,
		"Email":     user.Email,
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().AddDate(0, 0, 7).Unix(),
		"iat":       time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}
