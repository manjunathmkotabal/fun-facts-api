package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manjunath.kotabal/fun-facts-api/services"
)

func QueryByCategory(c *gin.Context) {
	category := c.Param("category")
	log.Printf("Querying trivia by category: %s", category)
	trivia, err := services.GetTriviaByCategory(category)
	if err != nil {
		log.Printf("Error retrieving trivia: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Retrieved trivia: %v", trivia)
	c.JSON(http.StatusOK, gin.H{"data": trivia})
}

func QueryTrivia(c *gin.Context) {
	trivia, err := services.GetRandomTrivia()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": trivia})
}
