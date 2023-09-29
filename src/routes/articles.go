package routes

import (
	"gin_golang/src/controllers"
	"gin_golang/src/middleware"

	"github.com/gin-gonic/gin"
)

func RouteArticle(articles *gin.RouterGroup) {
	articles.GET("/", middleware.IsAuth(), controllers.GetListArticle)
	articles.GET("/:slug", controllers.GetArticle)
	articles.POST("/", controllers.PostArticle)
	articles.PUT("/:id", controllers.UpdateArticle)
	articles.DELETE("/:id", controllers.DeleteArticle)
}
