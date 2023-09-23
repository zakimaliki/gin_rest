package routes

import (
	"gin_golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routing() {
	app := gin.Default()
	v1 := app.Group("/api/v1")
	{
		articles := v1.Group("/article")
		{
			articles.GET("/", controllers.GetHome)
			articles.GET("/:slug", controllers.GetArticle)
			articles.POST("/", controllers.PostArticle)
			articles.PUT("/:id", controllers.UpdateArticle)
			articles.DELETE("/:id", controllers.DeleteArticle)
		}
	}
	app.Run()
}
