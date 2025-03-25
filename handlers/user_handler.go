package handlers

import (
	"go-fiber-mongo-crud/models"
	"go-fiber-mongo-crud/services"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	createdUser, err := services.CreateUserService(user)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(201).JSON(createdUser)
}

func GetUsers(c *fiber.Ctx) error {
	users, err := services.GetUsersService()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := services.GetUserByIDService(id)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	updatedUser, err := services.UpdateUserService(id, user)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON(updatedUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := services.DeleteUserService(id); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.SendStatus(204)
}
