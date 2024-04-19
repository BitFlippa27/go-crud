//GET REQUEST WITH MONGODB
//controller
type UserController struct {
	userservice services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		userservice: userservice
	}
}
//working with struct and JSON
func (uc *UserController) GetUser(ctx *gin.Context) {
	username :=  ctx.Param("name")
	user, err := uc.GetUser(username)
	ctx.JSON(http.StatusOK, user)
}

//Service 
type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx context.Context
}
//working with bson and struct
func (us *UserServiceImpl) GetUser(name string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "username", Value: name}}
	err := us.usercollection.FindOne(us.ctx, query).Decode(&user)
	return user, err
}


//string -> bson -> struct -> JSON

