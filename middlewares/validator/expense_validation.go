package validator

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	. "github.com/wallet-tracky/Golang-backend/dto"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/util"
	"net/http"
	"time"
)

type ExpenseValidator struct {
}

func (validator ExpenseValidator) ValidateAddExpenseRequest(context *gin.Context) {

	var newExpense request.Expense
	err := context.BindJSON(&newExpense)
	if err != nil {
		logger.Error("Bad request",err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Status: 400,Message: err.Error()})
		return
	}

	date, err := time.Parse(util.TIME_LAYOUT, newExpense.Date)

	if err!=nil{
		logger.Info("Error while parsing date. Given date: ",newExpense.Date)
		logger.Error(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Status: 400,Message: "Invalid date format. Please add date as "+util.TIME_LAYOUT})
		return
	}

	if date.After(time.Now()){
		logger.Info("Spend date can't be a date in future. Given date: ",newExpense.Date)
		context.AbortWithStatusJSON(http.StatusBadRequest, "Spend date can't be a date in future. Given date: \",newExpense.Date")
		return
	}

	newExpense.Date=date.Format(util.TIME_LAYOUT)

	context.Set("request", &newExpense)
	context.Next()
}

func (validator ExpenseValidator) ValidateGetExpensesOfUser(context *gin.Context) {

	context.Next()
}
