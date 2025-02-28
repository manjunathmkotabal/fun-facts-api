package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjunath.kotabal/fun-facts-api/controllers"
	"github.com/manjunath.kotabal/fun-facts-api/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Apply rate limiter middleware to all routes
	r.Use(middleware.RateLimiter(1, 5)) // 1 request per second with a burst of 5

	// Public routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/refresh", controllers.RefreshToken)

	// Apply authentication middleware to protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/trivia/random", controllers.GetRandomTrivia)
		protected.GET("/trivia/category/:category", controllers.GetTriviaByCategory)
	}

	return r
}
