package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markDoesany/QuickyKanban/internal/config"
	"github.com/markDoesany/QuickyKanban/internal/models"
	"github.com/markDoesany/QuickyKanban/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{}, &models.Comment{})
	fmt.Println("Migration Completed")

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	routes.RegisterRoutes(r)
	r.Run(":8080")
}
