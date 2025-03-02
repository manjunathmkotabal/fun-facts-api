package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjunath.kotabal/fun-facts-api/controllers"
	"github.com/manjunath.kotabal/fun-facts-api/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/refresh", controllers.RefreshToken)

	// Admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware("admin"))
	{
		admin.GET("/query-by-category/:category", controllers.QueryByCategory)
		admin.GET("/query-trivia", controllers.QueryTrivia)
	}

	// User routes
	user := r.Group("/user")
	user.Use(middleware.AuthMiddleware("user"))
	{
		user.GET("/query-trivia", controllers.QueryTrivia)
	}

	return r
}
