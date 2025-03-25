package main

import (
	"go-fiber-mongo-crud/config"
	"go-fiber-mongo-crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load()

	// Initialize MongoDB Connection
	config.ConnectDatabase()

	app := fiber.New()

	// Register routes
	routes.RegisterRoutes(app)

	// Start server
	app.Listen(":3000")
}
