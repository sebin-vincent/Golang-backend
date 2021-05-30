package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/dto"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/service"
	"github.com/wallet-tracky/Golang-backend/util"
	"net/http"
	"strconv"
)

type ExpenseController struct {
	expenseService service.ExpenseService
}

func NewExpenseController(expenseService service.ExpenseService) *ExpenseController {
	return &ExpenseController{expenseService: expenseService}
}

func (controller *ExpenseController) AddExpense(ctx *gin.Context) {

	idObject, _ := ctx.Get("userId")

	userId:=idObject.(int)

	requestBody, _ := ctx.Get("expense")

	newExpense := requestBody.(*request.Expense)

	responseDTO, err := controller.expenseService.Save(newExpense,userId)

	if util.IsError(err) {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Something went wrong!"})
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
