package main

import (
	"events-rest-api/db"
	"events-rest-api/models"
	"events-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	setup()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}

func setup() {
	db.DB.AutoMigrate(&models.Event{}, &models.User{})
}
