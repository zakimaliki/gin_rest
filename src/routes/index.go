package routes

import (
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
		v1.Group("/article")
		{
			RouteArticle(v1)
		}
	}
	app.Run(os.Getenv("PORT"))
}
