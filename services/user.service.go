package services

import (
	"github.com/bitflippa27/go-crud/models"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(string) error
	InitialDataLoad() ([]*models.User, error)
}
