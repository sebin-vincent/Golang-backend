package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wallet-tracky/Golang-backend/src/dto"
	"github.com/wallet-tracky/Golang-backend/src/dto/request"
	"github.com/wallet-tracky/Golang-backend/src/util"
	"net/http"
	"strconv"
)

type ExpenseValidator struct {
}

func (validator ExpenseValidator) ValidateAddExpenseRequest(context *gin.Context) {

	var newExpense request.Expense
	err := context.BindJSON(&newExpense)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		context.Abort()
	}

	context.Set("expense", &newExpense)
	context.Next()
}

func (validator ExpenseValidator) ValidateGetExpensesOfUser(context *gin.Context) {

	_, err := strconv.Atoi(context.GetHeader("userId"))

	if util.IsError(err) {
		fmt.Println("Invalid userId")
		context.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid value for userId"})
	}

	context.Next()
}
