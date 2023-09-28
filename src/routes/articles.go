package routes

import (
	"gin_golang/src/controllers"
	"gin_golang/src/middleware"

	"github.com/gin-gonic/gin"
)

func Routing() {
	app := gin.Default()
	v1 := app.Group("/api/v1")
	{
		v1.GET("/auth/:provider", controllers.RedirectHandler)
		v1.GET("/auth/:provider/callback", controllers.CallbackHandler)

		articles := v1.Group("/article")

		{
			articles.GET("/", middleware.IsAuth(), controllers.GetListArticle)
			articles.GET("/:slug", controllers.GetArticle)
			articles.POST("/", controllers.PostArticle)
			articles.PUT("/:id", controllers.UpdateArticle)
			articles.DELETE("/:id", controllers.DeleteArticle)
		}
	}
	app.Run()
}
