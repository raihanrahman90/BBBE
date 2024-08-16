// routes/routes.go
package routes

import (
	"os"
	"rumahbelajar/handlers/article"
	"rumahbelajar/handlers/class"
	landingpage "rumahbelajar/handlers/landingPage"
	"rumahbelajar/handlers/teacher"
	"rumahbelajar/handlers/testimoni"
	"rumahbelajar/handlers/user"
	"rumahbelajar/middleware"

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
	r.GET("/fields", teacher.GetAllFields)
	r.GET("/class", class.GetClass)
	r.GET("/teachers", teacher.GetTeacher)
	r.GET("/article", article.GetArticle)
	r.GET("/testimoni", testimoni.GetTestimoni)
	r.GET("/article/:title", article.GetArticleByTitle)
	r.GET("/landing-page", landingpage.GetLandingPageData)

	protected := r.Group("/admin")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/teachers", teacher.GetTeacher)
		protected.POST("/teachers", teacher.CreateTeacher)
		protected.GET("/teachers/:id", teacher.GetTeacherById)
		protected.POST("/teachers/:id", teacher.UpdateTeacher)
		protected.DELETE("/teachers/:id", teacher.DeleteTeacher)

		protected.GET("/class", class.GetClass)
		protected.POST("/class", class.CreateClass)
		protected.GET("/class/:id", class.GetClassById)
		protected.POST("/class/:id", class.UpdateClass)
		protected.DELETE("/class/:id", class.DeleteClass)

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

		protected.GET("/landing-page", landingpage.GetLandingPageData)
		protected.POST("/landing-page", landingpage.UpdateDataLandingPage)
	}

	return r
}
