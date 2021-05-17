package router

import (
	"github.com/gin-gonic/gin"
	controller2 "github.com/wallet-tracky/Golang-backend/controller"
	validator2 "github.com/wallet-tracky/Golang-backend/middlewares/validator"
	service2 "github.com/wallet-tracky/Golang-backend/service"
)

type ExpenseRouter struct {
	controller *controller2.ExpenseController
}

func (router *ExpenseRouter) InitializeExpenseRouting(server *gin.Engine) {
	expenseValidator := validator2.ExpenseValidator{}

	server.POST("/expenses", expenseValidator.ValidateAddExpenseRequest, router.controller.AddExpense)

	server.GET("/expenses", expenseValidator.ValidateGetExpensesOfUser, router.controller.GetExpenses)
}

func NewExpenseRouter() *ExpenseRouter {
	videoService := service2.New()
	videoController := controller2.New(videoService)

	return &ExpenseRouter{
		controller: videoController,
	}
}
