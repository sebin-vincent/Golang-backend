package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wallet-tracky/Golang-backend/controller"
	. "github.com/wallet-tracky/Golang-backend/middlewares"
	"github.com/wallet-tracky/Golang-backend/middlewares/validator"
	"github.com/wallet-tracky/Golang-backend/repository"
	"github.com/wallet-tracky/Golang-backend/service"
)

type ExpenseRouter struct {
	controller *controller.ExpenseController
}

func (router *ExpenseRouter) InitializeExpenseRouting(server *gin.Engine) {
	expenseValidator := validator.ExpenseValidator{}

	server.POST("/expenses", Authenticate(""),expenseValidator.ValidateAddExpenseRequest, router.controller.AddExpense)

	server.GET("/expenses", expenseValidator.ValidateGetExpensesOfUser, router.controller.GetExpenses)
}

func NewExpenseRouter() *ExpenseRouter {
	expenseRepository := repository.NewExpenseRepository()
	expenseService := service.NewExpenseService(expenseRepository)
	expenseController := controller.NewExpenseController(expenseService)

	return &ExpenseRouter{
		controller: expenseController,
	}
}
