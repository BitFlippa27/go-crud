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
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

// Constructor returns instance of UserService
func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

// Receiver function == method of UserServiceImpl "class"
func (u *UserServiceImpl) GetUser(name string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "username", Value: name}} //db.collection.find({name: "elliot"})
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user *models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	if err != nil {
		return err
	}
	return err
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filterquery := bson.D{bson.E{Key: "username", Value: user.Name}}
	updatequery := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "username", Value: user.Name},
			bson.E{Key: "userage", Value: user.Age},
			bson.E{Key: "useraddress", Value: user.Address},
		}}}
	result, err := u.usercollection.UpdateOne(u.ctx, filterquery, updatequery)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	if err != nil {
		return errors.New("error in updateUser")
	}

	return nil
}

func (u *UserServiceImpl) DeleteUser(name string) error {
	filterquery := bson.D{bson.E{Key: "username", Value: name}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filterquery)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for deletion")
	}
	return nil
}

func (u *UserServiceImpl) InitialDataLoad() ([]*models.User, error) {
	var users []*models.User
	response, err := http.Get("https://jsonplaceholder.typicode.com/users")
	fmt.Println(response)
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

	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	var docs []interface{}
	for _, user := range users {
		docs = append(docs, user)
	}

	_, err = u.usercollection.InsertMany(u.ctx, docs)
	if err != nil {
		return nil, err
	}

	return users, nil
}
