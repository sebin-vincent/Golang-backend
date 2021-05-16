package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wallet-tracky/Golang-backend/src/controller"
	"github.com/wallet-tracky/Golang-backend/src/middlewares/validator"
	"github.com/wallet-tracky/Golang-backend/src/service"
)

type ExpenseRouter struct {
	controller *controller.ExpenseController
}

func (router *ExpenseRouter) InitializeExpenseRouting(server *gin.Engine) {
	expenseValidator := validator.ExpenseValidator{}

	server.POST("/expenses", expenseValidator.ValidateAddExpenseRequest, router.controller.AddExpense)

	server.GET("/expenses", expenseValidator.ValidateGetExpensesOfUser, router.controller.GetExpenses)
}

func NewExpenseRouter() *ExpenseRouter {
	videoService := service.New()
	videoController := controller.New(videoService)

	return &ExpenseRouter{
		controller: videoController,
	}
}
