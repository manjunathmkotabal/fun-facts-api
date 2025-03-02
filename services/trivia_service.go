package services

import (
	"github.com/manjunath.kotabal/fun-facts-api/models"
)

func GetRandomTrivia() ([]models.Trivia, error) {
	var trivia []models.Trivia
	if err := models.DB.Order("RANDOM()").Limit(1).Find(&trivia).Error; err != nil {
		return nil, err
	}
	return trivia, nil
}

func GetTriviaByCategory(category string) ([]models.Trivia, error) {
	var trivia []models.Trivia
	if err := models.DB.Where("category = ?", category).Find(&trivia).Error; err != nil {
		return nil, err
	}
	return trivia, nil
}
