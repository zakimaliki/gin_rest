package routes

import (
	"gin_golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func RouteUser(users *gin.RouterGroup) {
	users.GET("/auth/:provider", controllers.RedirectHandler)
	users.GET("/auth/:provider/callback", controllers.CallbackHandler)
}
