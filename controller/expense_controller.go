package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	dto2 "github.com/wallet-tracky/Golang-backend/dto"
	request2 "github.com/wallet-tracky/Golang-backend/dto/request"
	service2 "github.com/wallet-tracky/Golang-backend/service"
	util2 "github.com/wallet-tracky/Golang-backend/util"
	"net/http"
	"strconv"
)

type ExpenseController struct {
	expenseService service2.ExpenseService
}

func New(expenseService service2.ExpenseService) *ExpenseController {
	return &ExpenseController{expenseService: expenseService}
}

func (controller *ExpenseController) AddExpense(ctx *gin.Context) {

	requestBody, _ := ctx.Get("expense")

	newExpense := requestBody.(*request2.Expense)

	responseDTO, err := controller.expenseService.Save(newExpense)

	if util2.IsError(err) {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto2.ErrorResponse{Message: "Something went wrong!"})
		return
	}

	ctx.JSON(http.StatusOK, responseDTO)
}

func (controller *ExpenseController) GetExpenses(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.GetHeader("userId"))

	logger.Infof("Request to get expenses for user: %d", userId)
	userExpenses := controller.expenseService.FindAllExpenseOfUser(userId)

	ctx.JSON(http.StatusOK, userExpenses)
}
