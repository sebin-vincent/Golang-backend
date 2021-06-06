package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wallet-tracky/Golang-backend/controller"
	"github.com/wallet-tracky/Golang-backend/middlewares/validator"
	"github.com/wallet-tracky/Golang-backend/repository"
	"github.com/wallet-tracky/Golang-backend/service"
)

type UserRouter struct {
	userController controller.UserController
}

func (router *UserRouter) InitializeUserRouter(server *gin.Engine) {
	userValidator := validator.UserValidator{}

	server.POST("/user/signup", userValidator.ValidateUserSignUp, router.userController.UserSignUp)
	server.POST("/login",userValidator.LoginValidator,router.userController.Login)
}




func NewUserRouter() UserRouter{

	repo:=repository.NewUserRepository()
	userService:=service.NewUserService(repo)
	userController:=controller.NewUserController(userService)

	return UserRouter{userController: userController}
}