package main

import (
	"go-fiber-mongo-crud/config"
	"go-fiber-mongo-crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	config.ConnectDatabase()

	app := fiber.New()

	routes.RegisterRoutes(app)

	app.Listen(":3000")
}
