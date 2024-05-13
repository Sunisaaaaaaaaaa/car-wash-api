package main

import (
	"booking-api/config"
	"booking-api/routes"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config.InitDB()

	r := gin.Default()
	r.Use(cors.Default())
	routes.Routes(r)
	r.Run("localhost:8080")
}
