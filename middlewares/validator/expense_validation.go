package validator

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	. "github.com/wallet-tracky/Golang-backend/dto"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/util"
	"net/http"
	"regexp"
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

	date, err := time.Parse(util.DateTimeLayout, newExpense.Date)

	if err!=nil{
		logger.Info("Error while parsing date. Given date: ",newExpense.Date)
		logger.Error(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Status: 400,Message: "Invalid date format. Please add date as "+util.DateTimeLayout})
		return
	}

	if date.After(time.Now()){
		logger.Info("Spend date can't be a date in future. Given date: ",newExpense.Date)
		context.AbortWithStatusJSON(http.StatusBadRequest,  ErrorResponse{Status: 400,Message: "Spend date can't be a date in future. Given date: "+newExpense.Date})
		return
	}

	newExpense.Date=date.Format(util.DateTimeLayout)

	context.Set("request", &newExpense)
	context.Next()
}

func (validator ExpenseValidator) ValidateGetExpensesOfUser(context *gin.Context) {

	integerChecker := regexp.MustCompile("^\\d*$")


	if !(integerChecker.MatchString(context.DefaultQuery("offset","0"))&&
		integerChecker.MatchString(context.DefaultQuery("limit","5"))){

		//if either offset or limit are not positive integers (eg. alphabets and negative numbers) throw error
		context.AbortWithStatusJSON(400,ErrorResponse{Status: 400,Message: "Invalid arguments for pagination"})
		return
	}

	queryMap := context.Request.URL.Query()

	fromDate:=queryMap.Get("from")

	_, err:=time.Parse(util.DateLayout,fromDate)

	if err!=nil {
		logger.Info("Error while parsing 'from' date")
		logger.Info("Error",err.Error())
		context.AbortWithStatusJSON(400,ErrorResponse{Status: 400,Message: "Error while parsing 'from' date"})
		return
	}

	toDate:=queryMap.Get("to")

	_, err=time.Parse(util.DateLayout,toDate)

	if err!=nil {
		logger.Info("Error while parsing 'to' date")
		logger.Info("Error",err.Error())
		context.AbortWithStatusJSON(400,ErrorResponse{Status: 400,Message: "Error while parsing 'to' date"})
		return
	}

	context.Next()
}
