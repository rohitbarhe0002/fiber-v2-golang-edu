package middleware

import (
	"go-fiber-mongo-crud/utils"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware is a middleware that validates JWT tokens
func AuthMiddleware(c *fiber.Ctx) error {
	// Get token from the request header
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized, token missing"})
	}

	// Remove "Bearer " prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// Validate the JWT token
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Store user ID in Locals for use in handlers
	c.Locals("user_id", claims["user_id"])

	// Proceed to the next handler
	return c.Next()
}
