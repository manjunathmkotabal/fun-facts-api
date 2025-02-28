package services

import (
	"errors"
	"math/rand"
	"time"

	"github.com/manjunath.kotabal/fun-facts-api/models"
)

func GetRandomTrivia() (models.Trivia, error) {
	var trivia models.Trivia
	var count int64
	models.DB.Model(&models.Trivia{}).Count(&count)
	if count == 0 {
		return trivia, errors.New("no trivia found")
	}
	rand.Seed(time.Now().UnixNano())
	offset := rand.Intn(int(count))
	models.DB.Offset(offset).First(&trivia)
	return trivia, nil
}

func GetTriviaByCategory(category string) ([]models.Trivia, error) {
	var trivia []models.Trivia
	if err := models.DB.Where("category = ?", category).Find(&trivia).Error; err != nil {
		return trivia, err
	}
	return trivia, nil
}
