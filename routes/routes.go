// routes/routes.go
package routes

import (
	"os"
	"bbbe/handlers/article"
	"bbbe/handlers/item"
	landingpage "bbbe/handlers/landingPage"
	"bbbe/handlers/testimoni"
	"bbbe/handlers/user"
	"bbbe/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	r.Static("/static", os.Getenv("PATH_STATIC"))
	// Define routes
	r.POST("/login", user.Login)
	r.GET("/auth/refresh", user.RefreshToken)
	r.POST("/users", user.CreateUser)
	r.GET("/article", article.GetArticle)
	r.GET("/testimoni", testimoni.GetTestimoni)
	r.GET("/article/:title", article.GetArticleByTitle)
	r.GET("/item", article.GetArticle)
	r.GET("/item/:id", article.GetArticleById)
	r.GET("/landing-page", landingpage.GetLandingPageData)

	protected := r.Group("/admin")
	protected.Use(middleware.AuthMiddleware())
	{

		protected.GET("/article", article.GetArticle)
		protected.POST("/article", article.CreateArticle)
		protected.GET("/article/:id", article.GetArticleById)
		protected.POST("/article/:id", article.UpdateArticle)
		protected.DELETE("/article/:id", article.DeleteArticle)

		protected.GET("/testimoni", testimoni.GetTestimoni)
		protected.POST("/testimoni", testimoni.CreateTestimoni)
		protected.GET("/testimoni/:id", testimoni.GetTestimoniById)
		protected.POST("/testimoni/:id", testimoni.UpdateTestimoni)
		protected.DELETE("/testimoni/:id", testimoni.DeleteTestimoni)

		protected.POST("/users", user.CreateUser)
		protected.GET("/users", user.GetUser)
		protected.POST("/users/me", user.UserUpdatePassword)
		protected.DELETE("/users/:id", user.DeleteUser)

		protected.GET("/item", item.GetItem)
		protected.POST("/item", item.CreateItem)
		protected.GET("/item/:id", item.GetItemById)
		protected.POST("/item/:id", item.UpdateItem)
		protected.DELETE("/item/:id", item.DeleteItem)
		protected.POST("/item-image", item.CreateItemImage)
		protected.DELETE("/item-image/:id", item.DeleteItemImage)

		protected.GET("/landing-page", landingpage.GetLandingPageData)
		protected.POST("/landing-page", landingpage.UpdateDataLandingPage)

		protected.GET("/transaction", transaction)
	}

}
