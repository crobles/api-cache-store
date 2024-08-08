package main

import (
	configFiber "api-cache-store/config/fiber"
	database "api-cache-store/database/redis"
	fiberModel "api-cache-store/internal/models/fiber"
	router "api-cache-store/router"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

var config fiberModel.Config = configFiber.SetConfigFiber()

func main() {
	fmt.Println("Starting server")
	time.Sleep(500 * time.Millisecond)
	app := fiber.New(fiber.Config{
		Prefork:	config.Prefork,
		CaseSensitive:	config.CaseSensitive,
		StrictRouting:  config.StrictRouting,
		ServerHeader:	config.ServerHeader,
		AppName:	config.AppName,
	})

	fmt.Println("Connecting to database")
	time.Sleep(500 * time.Millisecond)
	_, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to database with error: ", err)
		fmt.Println("Shutting down server")
		os.Exit(1)
	}

	fmt.Println("Configuring routes")
	time.Sleep(500 * time.Millisecond)
	router.SetupRoutes(app)

	fmt.Println("Starting server on port 8080")
	app.Listen(":8080")
}