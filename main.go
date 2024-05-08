package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bitflippa27/go-crud/controllers"
	"github.com/bitflippa27/go-crud/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	usercollection *mongo.Collection
	ctx            context.Context
	mongoclient    *mongo.Client
)

func init() {
	ctx = context.TODO()

	mongoconnection := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err := mongo.Connect(ctx, mongoconnection)
	if err != nil {
		log.Fatal()
	}
	err = mongoclient.Ping(ctx, &readpref.ReadPref{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo is up and running!")

	usercollection = mongoclient.Database("gocrud").Collection("users") //Models Zugriff
	userservice = services.NewUserService(usercollection, ctx)          //Services Zugriff
	usercontroller = controllers.NewUserController(userservice)         //Controllers Zugriff interface methoden
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx) // disconnect mongodb if app shuts down

	basepath := server.Group("/v1") // v1/user/create

	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":7777"))
}
