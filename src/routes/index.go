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
	app.MaxMultipartMemory = 1 // 8 MiB
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
		users := v1.Group("/auth")
		{
			users.GET("/auth/:provider", controllers.RedirectHandler)
			users.GET("/auth/:provider/callback", controllers.CallbackHandler)
		}
		articles := v1.Group("/article")
		{
			articles.GET("/", controllers.GetListArticle)
			articles.GET("/:slug", controllers.GetArticle)
			articles.POST("/", controllers.PostArticle)
			articles.PUT("/:slug", controllers.UpdateArticle)
			articles.DELETE("/:id", controllers.DeleteArticle)
			articles.POST("/upload", controllers.Uploadfile)
			articles.GET("/find", controllers.FindArticle)
			articles.GET("/paginate", controllers.PaginatSortArticle)
		}
		test := v1.Group("/test")
		{
			test.GET("/", controllers.TestList)
			test.GET("/:slug", controllers.Test)
			test.POST("/", controllers.TestPost)
			test.PUT("/:id", controllers.TestUpdate)
			test.DELETE("/:id", controllers.TestDelete)
		}
	}
	app.Run(os.Getenv("PORT"))
}
