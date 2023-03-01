package userserv

import (
	"client-rest/models"
)

type UserService interface {
	FindUserById(id string) (*models.UserResponse, error)
	GetAllUsers() ([]*models.UserResponse, error)
	CreateUser(*models.UserDB) (*models.UserResponse, error)
	UpdateUser(id string, user *models.UserUpdate) (*models.UserResponse, error)
	DeleteUser(string) error
	FindUserByEmail(email string) (*models.UserResponse, error)
}
