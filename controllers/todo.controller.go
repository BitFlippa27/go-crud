package controllers

import (
	"fmt"
	"net/http"

	"github.com/bitflippa27/go-crud/models"
	"github.com/bitflippa27/go-crud/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoController struct {
	todoservice services.TodoService
}

// Constructor returns instance of UserController
// injecting UserService dependency into UserController
func NewTodoController(todoservice services.TodoService) TodoController {
	return TodoController{
		todoservice: todoservice,
	}
}

/*
	func (t *TodoController) GetTodo(ctx *gin.Context) {
		username := ctx.Param("name")
		user, err := uc.userservice.GetUser(username)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) // Error when retrieving from DB
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
*/
func (tc *TodoController) CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil { //json body from incoming HTTP Request into struct (user variable)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) //sending 400 code in response json body back
		return
	}
	err := tc.todoservice.CreateTodo(&todo) //invoking service
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()}) // Error when saving to mongoDB
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (tc *TodoController) GetAllTodos(ctx *gin.Context) {
	fmt.Printf("ctx.Writer: %v\n", ctx.Writer)
	todos, err := tc.todoservice.GetAllTodos()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (tc *TodoController) DeleteTodo(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	err = tc.todoservice.DeleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (tc *TodoController) UpdateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil { //JSON to Go data structure
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.todoservice.UpdateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"}) // Go into JSON
}

func (tc *TodoController) MarkCompleted(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil { //JSON to Go data structure
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.todoservice.UpdateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"}) // Go into JSON
}

func (tc *TodoController) InitialDataLoad(ctx *gin.Context) {
	todos, err := tc.todoservice.InitialDataLoad()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todos)

}

func (tc *TodoController) RegisterUserRoutes(rg *gin.RouterGroup) {
	todoroute := rg.Group("/todos") //base path = home route
	todoroute.POST("/create", tc.CreateTodo)
	//userroute.GET("/get/:name", uc.GetUser)
	todoroute.GET("/getall", tc.GetAllTodos)
	todoroute.GET("/initial", tc.InitialDataLoad)
	todoroute.DELETE("/delete/:id", tc.DeleteTodo)
	todoroute.PATCH("/update", tc.UpdateTodo)
	todoroute.PATCH("/:id/complete", tc.UpdateTodo)
}
