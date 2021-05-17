package service

import (
	logger "github.com/sirupsen/logrus"
	request2 "github.com/wallet-tracky/Golang-backend/dto/request"
	response2 "github.com/wallet-tracky/Golang-backend/dto/response"
	model2 "github.com/wallet-tracky/Golang-backend/model"
	repository2 "github.com/wallet-tracky/Golang-backend/repository"
)

type ExpenseService interface {
	Save(expense *request2.Expense) (*response2.Expense, error)
	FindAllExpenseOfUser(userId int) []response2.Expense
}

type expenseService struct {
	expenseRepository repository2.ExpenseRepository
}

func New() ExpenseService {
	return &expenseService{expenseRepository: repository2.NewExpenseRepository()}
}

func (expenseService *expenseService) Save(expense *request2.Expense) (*response2.Expense, error) {

	var responseDTO *response2.Expense

	newExpense := makeNewExpenseEntity(expense) //private method call to get new expense model from model.Expense

	err := expenseService.expenseRepository.Save(newExpense)

	responseDTO = makeNewExpenseResponseDTO(newExpense) //private method call to get responseDTO from model.Expense
	return responseDTO, err
}

func (expenseService *expenseService) FindAllExpenseOfUser(userId int) []response2.Expense {

	logger.Info("Get user expenses")

	expenses := expenseService.expenseRepository.FindByUserId(userId)

	userExpenses := make([]response2.Expense, len(*expenses))

	var expenseDTO *response2.Expense
	for index, expense := range *expenses {
		expenseDTO = makeNewExpenseResponseDTO(&expense)
		userExpenses[index] = *expenseDTO
	}

	return userExpenses
}


func makeNewExpenseEntity(expense *request2.Expense) *model2.Expense {

	newExpense := &model2.Expense{}
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

func makeNewExpenseResponseDTO(expense *model2.Expense) *response2.Expense {

	expenseResponse := new(response2.Expense)

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
