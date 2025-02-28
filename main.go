package main

import (
	"github.com/manjunath.kotabal/fun-facts-api/config"
	"github.com/manjunath.kotabal/fun-facts-api/models"
	"github.com/manjunath.kotabal/fun-facts-api/routes"
)

func main() {
	config.LoadConfig()
	models.InitDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
