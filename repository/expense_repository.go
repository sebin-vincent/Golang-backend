package repository

import (
	"fmt"
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/util"
	"gorm.io/gorm"
)


type ExpenseRepository interface {
	Save(expense *model.Expense) error
	FindByUserId(userId int) *[]model.Expense
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
		fmt.Println(transaction.Error.Error())
	}

	return &expenses
}

func NewExpenseRepository() ExpenseRepository {

	return &expenseRepository{database: util.DB}
}
