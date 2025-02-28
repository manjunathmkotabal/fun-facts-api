package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manjunath.kotabal/fun-facts-api/services"
)

func GetRandomTrivia(c *gin.Context) {
	trivia, err := services.GetRandomTrivia()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trivia)
}

func GetTriviaByCategory(c *gin.Context) {
	category := c.Param("category")
	trivia, err := services.GetTriviaByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trivia)
}
