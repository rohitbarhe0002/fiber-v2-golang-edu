package services

import (
	"go-fiber-mongo-crud/models"
	"go-fiber-mongo-crud/repositories"
)

func CreateUserService(user models.User) (*models.User, error) {
	return repositories.CreateUser(user)
}

func GetUsersService() ([]models.User, error) {
	return repositories.GetUsers()
}

func GetUserByIDService(id string) (*models.User, error) {
	return repositories.GetUserByID(id)
}

func UpdateUserService(id string, user models.User) (*models.User, error) {
	return repositories.UpdateUser(id, user)
}

func DeleteUserService(id string) error {
	return repositories.DeleteUser(id)
}
