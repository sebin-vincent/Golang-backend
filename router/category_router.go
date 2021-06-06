package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wallet-tracky/Golang-backend/controller"
	"github.com/wallet-tracky/Golang-backend/middlewares"
	"github.com/wallet-tracky/Golang-backend/repository"
	"github.com/wallet-tracky/Golang-backend/service"
)


type CategoryRouter struct{

	controller controller.CategoryController
}

func (router *CategoryRouter) InitializeCategoryRouting(server *gin.Engine){

	server.GET("/categories/:type",middlewares.Authenticate(""),router.controller.GetCategories)
}


func NewCategoryRouter() CategoryRouter{

	repo:=repository.NewCategoryRepository();
	service:= service.NewCategoryService(repo)
	controller:= controller.NewCategoryController(service)

	return CategoryRouter{controller: controller}
}