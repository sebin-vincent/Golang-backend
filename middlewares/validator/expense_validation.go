package validator

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	dto2 "github.com/wallet-tracky/Golang-backend/dto"
	request2 "github.com/wallet-tracky/Golang-backend/dto/request"
	util2 "github.com/wallet-tracky/Golang-backend/util"
	"net/http"
	"strconv"
)

type ExpenseValidator struct {
}

func (validator ExpenseValidator) ValidateAddExpenseRequest(context *gin.Context) {

	var newExpense request2.Expense
	err := context.BindJSON(&newExpense)
	if err != nil {
		logger.Error(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		context.Abort()
	}

	context.Set("expense", &newExpense)
	context.Next()
}

func (validator ExpenseValidator) ValidateGetExpensesOfUser(context *gin.Context) {

	_, err := strconv.Atoi(context.GetHeader("userId"))

	if util2.IsError(err) {
		logger.Error("Invalid userId")
		context.AbortWithStatusJSON(http.StatusBadRequest, dto2.ErrorResponse{Message: "Invalid value for userId"})
	}

	context.Next()
}
