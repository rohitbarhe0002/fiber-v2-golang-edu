package routes

import (
	"go-fiber-mongo-crud/handlers"
	"go-fiber-mongo-crud/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Public routes
	api.Post("/auth/register", handlers.RegisterUser)
	api.Post("/auth/login", handlers.LoginUser)

	// Protected routes (requires authentication)
	api.Use(middleware.AuthMiddleware)

	api.Post("/users", handlers.CreateUser)
	api.Get("/users", handlers.GetUsers)
	api.Get("/users/:id", handlers.GetUserByID)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.DeleteUser)
}
