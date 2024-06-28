package main

import (
	database "api-cache-store/database/redis"
	router "api-cache-store/router"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

var mamalo string = "mamalo"

func main() {
	fmt.Println("Starting server")
	time.Sleep(1 * time.Second)
	app := fiber.New()

	fmt.Println("Connecting to database")
	time.Sleep(1 * time.Second)
	_, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to database with error: ", err)
		fmt.Println("Shutting down server")
		os.Exit(1)
	}

	fmt.Println("Configuring routes")
	time.Sleep(1 * time.Second)
	router.SetupRoutes(app)

	fmt.Println("Starting server on port 8080")
	app.Listen(":8080")
}