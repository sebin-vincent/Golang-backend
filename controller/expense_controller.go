package controller

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/service"
	"net/http"
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

	requestBody, _ := ctx.Get("request")

	newExpense := requestBody.(*request.Expense)

	responseDTO, err := controller.expenseService.Save(newExpense,userId)

	if err!=nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,err)
		return
	}

	ctx.JSON(http.StatusOK, responseDTO)
}

func (controller *ExpenseController) GetExpenses(ctx *gin.Context) {
	id, _ := ctx.Get("userId")
	userId:=id.(int)

	logger.Infof("Request to get expenses for user: %d", userId)
	userExpenses := controller.expenseService.FindAllExpenseOfUser(userId)

	ctx.JSON(http.StatusOK, userExpenses)
}
