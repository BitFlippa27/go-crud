package services

import (
	"github.com/bitflippa27/go-crud/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService interface {
	CreateTodo(*models.Todo) error
	//GetTodo(string) (*models.Todo, error)
	GetAllTodos() ([]*models.Todo, error)
	UpdateTodo(*models.Todo) error
	MarkCompleted(*models.Todo) error
	DeleteTodo(primitive.ObjectID) error
	InitialDataLoad() ([]*models.Todo, error)
}
