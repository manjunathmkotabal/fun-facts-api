package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjunath.kotabal/fun-facts-api/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/trivia/random", controllers.GetRandomTrivia)
	r.GET("/trivia/category/:category", controllers.GetTriviaByCategory)

	return r
}
