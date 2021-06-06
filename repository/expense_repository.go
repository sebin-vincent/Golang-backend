package repository

import (
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/dto"
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/util"
	"gorm.io/gorm"
	"time"
)


type ExpenseRepository interface {
	Save(expense *model.Expense) error
	FindByUserId(userId int) *[]model.Expense
	FindByUserIdAndSpendDateBetween(userId int, from time.Time, to time.Time,pageable dto.Pageable ) *[]model.Expense
	FindByUserIdAndCategoriesAndSpendDateBetween(
		userId int,
		categories []string,
		from time.Time,
		to time.Time,
		pageable dto.Pageable,
	) *[]model.Expense
}

type expenseRepository struct {
	database *gorm.DB
}

func (repository *expenseRepository) Save(expense *model.Expense) error {

	createdExpense := repository.database.Create(expense)
	err := createdExpense.Error

	return err
}

func (repository *expenseRepository) FindByUserId(userId int) *[]model.Expense {

	var expenses []model.Expense

	transaction := util.DB.Find(&expenses, map[string]interface{}{"userId": userId})

	if util.IsError(transaction.Error) {
		logger.Error(transaction.Error.Error())
	}

	return &expenses
}

func (repository *expenseRepository) FindByUserIdAndSpendDateBetween(
	userId int,
	from time.Time,
	to time.Time,
	pageable dto.Pageable,
	) *[]model.Expense{

	var expenses []model.Expense

	transaction := repository.database.
		Limit(pageable.Limit).
		Offset(pageable.OffSet).
		Find(&expenses, "userId=? AND spend_date between ? AND ?", userId, from, to)

	if transaction.Error!=nil {
		logger.Error(transaction.Error.Error())
		return nil
	}

	return &expenses
}

func (repository *expenseRepository) FindByUserIdAndCategoriesAndSpendDateBetween(
	userId int,
	categories []string,
	from time.Time,
	to time.Time,
	pageable dto.Pageable,
) *[]model.Expense{

	var expenses []model.Expense

	transaction := repository.database.
		Limit(pageable.Limit).
		Offset(pageable.OffSet).
		Find(&expenses, "userId=? AND category in ? AND spend_date between ? AND ?", userId, categories,from, to)

	if transaction.Error!=nil {
		logger.Error(transaction.Error.Error())
		return nil
	}

	return &expenses
}


func NewExpenseRepository() ExpenseRepository {

	return &expenseRepository{database: util.DB}
}
