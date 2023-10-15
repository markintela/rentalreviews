package controllers

import (
	services "rentalreviewspt/services/services"
	types "rentalreviewspt/services/types"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	var user types.User
	ctx.ShouldBindJSON(&user)
	ctx.JSON(200, "Ok")
}

func (uc *UserController) GetUser(ctx *gin.Context) {

	ctx.JSON(200, "Ok")
}

func (uc *UserController) GetUserByName(ctx *gin.Context) {

	ctx.JSON(200, "Ok")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {

	ctx.JSON(200, "Ok")
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {

	ctx.JSON(200, "Ok")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

	ctx.JSON(200, "Ok")
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {

	userRoute := rg.Group("user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:name", uc.CreateUser)
	userRoute.GET("/getall", uc.CreateUser)
	userRoute.PATCH("/update", uc.CreateUser)
	userRoute.DELETE("/delete", uc.CreateUser)
}
