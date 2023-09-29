package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Router() {
	app := gin.Default()
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
