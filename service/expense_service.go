package service

import (
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/dto/response"
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/repository"
)

type ExpenseService interface {
	Save(expense *request.Expense) (*response.Expense, error)
	FindAllExpenseOfUser(userId int) []response.Expense
}

type expenseService struct {
	expenseRepository repository.ExpenseRepository
}

func NewExpenseService(expenseRepository repository.ExpenseRepository) ExpenseService {
	return &expenseService{expenseRepository: expenseRepository}
}

func (expenseService *expenseService) Save(expense *request.Expense) (*response.Expense, error) {

	var responseDTO *response.Expense

	newExpense := makeNewExpenseEntity(expense) //private method call to get new expense model from model.Expense

	err := expenseService.expenseRepository.Save(newExpense)

	responseDTO = makeNewExpenseResponseDTO(newExpense) //private method call to get responseDTO from model.Expense
	return responseDTO, err
}

func (expenseService *expenseService) FindAllExpenseOfUser(userId int) []response.Expense {

	logger.Info("Get user expenses")

	expenses := expenseService.expenseRepository.FindByUserId(userId)

	userExpenses := make([]response.Expense, len(*expenses))

	var expenseDTO *response.Expense
	for index, expense := range *expenses {
		expenseDTO = makeNewExpenseResponseDTO(&expense)
		userExpenses[index] = *expenseDTO
	}

	return userExpenses
}


func makeNewExpenseEntity(expense *request.Expense) *model.Expense {

	newExpense := &model.Expense{}
	newExpense.UserId=1
	newExpense.Description = expense.Description
	newExpense.Amount = expense.Amount
	newExpense.SpendFrom = expense.SpendFrom
	newExpense.Date = expense.Date
	newExpense.Category = expense.Category
	newExpense.AdditionalNotes = expense.AdditionalNotes
	newExpense.Image = expense.Image
	newExpense.IsCounted = expense.IsCounted
	newExpense.Tag = expense.Tag
	newExpense.AddedAs = expense.AddedAs
	newExpense.IsReviewed = expense.IsReviewed

	return newExpense
}

func makeNewExpenseResponseDTO(expense *model.Expense) *response.Expense {

	expenseResponse := new(response.Expense)

	expenseResponse.Id = expense.Id
	expenseResponse.Description = expense.Description
	expenseResponse.Amount = expense.Amount
	expenseResponse.SpendFrom = expense.SpendFrom
	expenseResponse.Date = expense.Date
	expenseResponse.Category = expense.Category
	expenseResponse.AdditionalNotes = expense.AdditionalNotes
	expenseResponse.Image = expense.Image
	expenseResponse.IsCounted = expense.IsCounted
	expenseResponse.Tag = expense.Tag
	expenseResponse.AddedAs = expense.AddedAs
	expenseResponse.IsReviewed = expense.IsReviewed

	return expenseResponse
}
