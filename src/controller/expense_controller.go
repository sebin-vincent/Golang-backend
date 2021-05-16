package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wallet-tracky/Golang-backend/src/dto"
	"github.com/wallet-tracky/Golang-backend/src/dto/request"
	"github.com/wallet-tracky/Golang-backend/src/service"
	"github.com/wallet-tracky/Golang-backend/src/util"
	"net/http"
	"strconv"
)

type ExpenseController struct {
	expenseService service.ExpenseService
}

func New(expenseService service.ExpenseService) *ExpenseController {
	return &ExpenseController{expenseService: expenseService}
}

func (controller *ExpenseController) AddExpense(ctx *gin.Context) {

	requestBody, _ := ctx.Get("expense")

	newExpense := requestBody.(*request.Expense)

	responseDTO, err := controller.expenseService.Save(newExpense)

	if util.IsError(err) {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Something went wrong!"})
		return
	}

	ctx.JSON(http.StatusOK, responseDTO)
}

func (controller *ExpenseController) GetExpenses(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.GetHeader("userId"))

	fmt.Printf("Request to get expenses for user: %d", userId)
	userExpenses := controller.expenseService.FindAllExpenseOfUser(userId)

	ctx.JSON(http.StatusOK, userExpenses)
}
