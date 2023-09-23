package main

import (
	"gin_golang/config"
	"gin_golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		articles := v1.Group("/article")
		{
			articles.GET("/", routes.GetHome)
			articles.GET("/:slug", routes.GetArticle)
			articles.POST("/", routes.PostArticle)
			articles.PUT("/:id", routes.UpdateArticle)
			articles.DELETE("/:id", routes.DeleteArticle)
		}
	}

	router.Run()
}
