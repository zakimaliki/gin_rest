package routes

import (
	"gin_golang/src/controllers"
	"os"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/dvwright/xss-mw"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Router() {
	app := gin.Default()
	app.Use(helmet.Default())
	var xssMdlwr xss.XssMw
	app.Use(xssMdlwr.RemoveXss())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	v1 := app.Group("/api/v1")
	{
		v1.Group("/auth")
		{
			RouteUser(v1)
		}
		articles := v1.Group("/article")
		{
			articles.GET("/", controllers.GetListArticle)
			articles.GET("/:slug", controllers.GetArticle)
			articles.POST("/", controllers.PostArticle)
			articles.PUT("/:id", controllers.UpdateArticle)
			articles.DELETE("/:id", controllers.DeleteArticle)
			articles.POST("/upload", controllers.Uploadfile)
			articles.GET("/find", controllers.FindArticle)
			articles.GET("/test", controllers.FindTest)

		}
	}
	app.Run(os.Getenv("PORT"))
}
