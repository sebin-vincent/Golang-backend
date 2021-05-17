package repository

import (
	"fmt"
	model2 "github.com/wallet-tracky/Golang-backend/model"
	util2 "github.com/wallet-tracky/Golang-backend/util"
	"gorm.io/gorm"
)

type ExpenseRepository interface {

	Save(expense *model2.Expense) error
	FindByUserId(userId int) *[]model2.Expense

}


type expenseRepository struct {
	database *gorm.DB
}

func (repository *expenseRepository) Save(expense *model2.Expense) error{

	createdExpense := repository.database.Create(expense)
	err := createdExpense.Error

	return err
}


func (repository *expenseRepository) FindByUserId(userId int) *[]model2.Expense {

	var expenses []model2.Expense


	transaction := util2.DB.Find(&expenses, map[string]interface{}{"userId": userId})

	if util2.IsError(transaction.Error){
		fmt.Println(transaction.Error.Error())
	}

	return &expenses
}


func NewExpenseRepository() ExpenseRepository {

	return &expenseRepository{database: util2.DB}
}
