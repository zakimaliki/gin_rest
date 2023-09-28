package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	return checkJWT()
}

func checkJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte(os.Getenv("JWT_SECRET_KEY")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println(claims["user_id"], claims["email"])
			} else {
				c.JSON(422, gin.H{
					"msg":   "Invalid Token",
					"error": err,
				})
				c.Abort()
				return
			}
		} else {
			c.JSON(422, gin.H{
				"msg": "Server Need Token",
			})
			c.Abort()
			return
		}

	}

}
