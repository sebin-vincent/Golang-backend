package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/service"
)



type UserController struct {
	userService service.UserService
}

func (controller *UserController) UserSignUp(context *gin.Context){

	requestBody,_:=context.Get("request")

	userRequestBody:=requestBody.(*request.UserSignUp)

	resp, err := controller.userService.UserSignup(userRequestBody)

	if err!=nil{
		context.AbortWithStatusJSON(err.Status,err)
		return
	}

	context.JSON(201,resp)
}

func (controller *UserController) Login(context *gin.Context){

	requestBody,_:=context.Get("request")
	loginRequest := requestBody.(*request.LoginRequest)

	responseBody, err := controller.userService.Login(*loginRequest)

	if err!=nil{
		context.AbortWithStatusJSON(err.Status,err.Message)
		return
	}

	context.JSON(200,responseBody)

}

func NewUserController(userService service.UserService) UserController{
	return UserController{userService: userService}
}