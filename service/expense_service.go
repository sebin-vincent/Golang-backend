package service

import (
	logger "github.com/sirupsen/logrus"
	. "github.com/wallet-tracky/Golang-backend/dto"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/dto/response"
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/repository"
	"time"
)

type ExpenseService interface {
	Save(expense *request.Expense, userId int) (*response.Expense, *ErrorResponse)
	FindAllExpenseOfUser(userId int,
		from time.Time,
		to time.Time,
		categories []string,
		pageable Pageable,
	) []response.Expense
}

type expenseService struct {
	expenseRepository repository.ExpenseRepository
}

func NewExpenseService(expenseRepository repository.ExpenseRepository) ExpenseService {
	return &expenseService{expenseRepository: expenseRepository}
}

func (expenseService *expenseService) Save(expense *request.Expense, userId int) (*response.Expense, *ErrorResponse) {

	var responseDTO *response.Expense

	newExpense := makeNewExpenseModel(expense, userId) //private method call to get new expense model from model.Expense

	err := expenseService.expenseRepository.Save(newExpense)

	if err != nil {
		logger.Error("Error while saving expense data. ", err)
		return nil, &ErrorResponse{Status: 500, Message: "Something went wrong!"}
	}

	responseDTO = makeNewExpenseResponseDTO(newExpense) //private method call to get responseDTO from model.Expense
	return responseDTO, nil
}

func (expenseService *expenseService) FindAllExpenseOfUser(
	userId int,
	from time.Time,
	to time.Time,
	categories []string,
	pageable Pageable) []response.Expense {
	logger.Debug("Inside get user expense service")

	var expenses *[]model.Expense

	to = to.Add(time.Hour * 24) //Add 24 hours to get expense inclusive for to date

	if len(categories) == 0 {
		logger.Info("Length of categories =", len(categories))
		expenses = expenseService.expenseRepository.FindByUserIdAndSpendDateBetween(userId, from, to,pageable)
	} else {
		logger.Info("Length of categories =", len(categories))
		expenses = expenseService.expenseRepository.
			FindByUserIdAndCategoriesAndSpendDateBetween(userId, categories, from, to,pageable)
	}

	userExpenses := make([]response.Expense, len(*expenses))

	var expenseDTO *response.Expense
	for index, expense := range *expenses {
		expenseDTO = makeNewExpenseResponseDTO(&expense)
		userExpenses[index] = *expenseDTO
	}

	return userExpenses
}

func makeNewExpenseModel(expense *request.Expense, userId int) *model.Expense {

	newExpense := &model.Expense{}
	newExpense.UserId = userId
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
