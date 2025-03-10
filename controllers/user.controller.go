package controllers

import (
	"net/http"

	"github.com/bitflippa27/go-crud/models"
	"github.com/bitflippa27/go-crud/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userservice services.UserService
}

// Constructor returns instance of UserController
// injecting UserService dependency into UserController
func NewUserController(userservice services.UserService) UserController {
	return UserController{
		userservice: userservice,
	}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.userservice.GetUser(username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) // Error when retrieving from DB
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil { //json body from incoming HTTP Request into struct (user variable)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) //sending 400 code in response json body back
		return
	}
	err := uc.userservice.CreateUser(&user) //saving to mongoDB
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) // Error when saving to mongoDB
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.userservice.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.userservice.DeleteUser(username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil { //JSON to Go data structure
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.userservice.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"}) // Go into JSON
}

func (uc *UserController) InitialDataLoad(ctx *gin.Context) {
	todos, err := uc.userservice.InitialDataLoad()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todos)

}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/users") //base path = home route
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAllUsers)
	userroute.GET("/initial", uc.InitialDataLoad)
	userroute.DELETE("/delete/:id", uc.DeleteUser)
	userroute.PATCH("/update", uc.UpdateUser)
}
