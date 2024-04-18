package controllers

import (
	"github.com/bitflippa27/go-crud/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userservice services.UserService
}

func New(userService services.UserService) UserController {
	return UserController{
		userservice: userService,
	}
}

func (uc UserController) GetUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group(("/user")) //base path = home route
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.DELETE("/delete", uc.DeleteUser)
	userroute.PATCH("/update", uc.UpdateUser)

}
