package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bitflippa27/go-crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoServiceImpl struct {
	todocollection *mongo.Collection
	ctx            context.Context
}

// Constructor returns instance of UserService
func NewTodoService(todocollection *mongo.Collection, ctx context.Context) TodoService {
	return &TodoServiceImpl{
		todocollection: todocollection,
		ctx:            ctx,
	}
}

/*
func (t *TodoServiceImpl) GetTodo(name string) (*models.Todo, error) {
	var todo *models.Todo
	query := bson.D{bson.E{Key: "username", Value: name}} //db.collection.find({name: "elliot"})
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, err
}
*/

func (t *TodoServiceImpl) GetAllTodos() ([]*models.Todo, error) {
	fmt.Println("GetAllTodos")
	var todos []*models.Todo
	cursor, err := t.todocollection.Find(t.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(t.ctx) {
		var todo *models.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(t.ctx)

	if len(todos) == 0 {
		return nil, errors.New("documents not found")
	}
	return todos, nil
}

func (t *TodoServiceImpl) CreateTodo(todo *models.Todo) error {
	_, err := t.todocollection.InsertOne(t.ctx, todo)
	if err != nil {
		return err
	}
	return err
}

func (t *TodoServiceImpl) UpdateTodo(todo *models.Todo) error {
	filterquery := bson.D{bson.E{Key: "_id", Value: todo.Id}}
	updatequery := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "title", Value: todo.Title},
		}}}
	result, err := t.todocollection.UpdateOne(t.ctx, filterquery, updatequery)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	if err != nil {
		return errors.New("error in updateTodo")
	}

	return nil
}

func (t *TodoServiceImpl) MarkCompleted(todo *models.Todo) error {
	filterquery := bson.D{bson.E{Key: "_id", Value: todo.Id}}
	updatequery := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "completed", Value: todo.Completed},
		}}}
	result, err := t.todocollection.UpdateOne(t.ctx, filterquery, updatequery)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	if err != nil {
		return errors.New("error in updateTodo")
	}

	return nil
}

func (t *TodoServiceImpl) DeleteTodo(id primitive.ObjectID) error {
	filterquery := bson.D{bson.E{Key: "_id", Value: id}}
	result, _ := t.todocollection.DeleteOne(t.ctx, filterquery)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for deletion")
	}
	return nil
}

func (t *TodoServiceImpl) InitialDataLoad() ([]*models.Todo, error) {
	var todos []*models.Todo
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println(response)
		return nil, fmt.Errorf("got HTTP status %d", response.StatusCode)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &todos)
	if err != nil {
		return nil, err
	}

	var docs []interface{}
	for _, todo := range todos {
		docs = append(docs, todo)
	}

	_, err = t.todocollection.InsertMany(t.ctx, docs)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
