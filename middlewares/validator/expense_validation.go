package validator

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/dto"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/util"
	"net/http"
	"strconv"
)

type ExpenseValidator struct {
}

func (validator ExpenseValidator) ValidateAddExpenseRequest(context *gin.Context) {

	var newExpense request.Expense
	err := context.BindJSON(&newExpense)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.Set("expense", &newExpense)
	context.Next()
}

func (validator ExpenseValidator) ValidateGetExpensesOfUser(context *gin.Context) {

	_, err := strconv.Atoi(context.GetHeader("userId"))

	if util.IsError(err) {
		logger.Error("Invalid userId")
		context.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid value for userId"})
		return
	}

	context.Next()
}
