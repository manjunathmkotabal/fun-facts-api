package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Trivia struct {
	gorm.Model
	Category string `json:"category"`
	Content  string `json:"content"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("trivia.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Trivia{}, &User{})

	// Seed data
	seedData()
}

func seedData() {
	trivia := []Trivia{
		{Category: "science", Content: "The Eiffel Tower can be 15 cm taller during the summer."},
		{Category: "history", Content: "The shortest war in history lasted 38 minutes."},
		{Category: "tech", Content: "The first computer virus was created in 1983."},
	}

	for _, t := range trivia {
		DB.FirstOrCreate(&t, Trivia{Content: t.Content})
	}
}
