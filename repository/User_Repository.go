package repository

import (
	"errors"
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/util"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type userRepository struct {
	database *gorm.DB
}

func (repo *userRepository) Save(user *model.User) error {
	createdExpense := repo.database.Create(user)
	err := createdExpense.Error

	return err
}

func (repo *userRepository) FindByEmail(email string) (*model.User, error) {

	var user model.User

	err := repo.database.First(&user, map[string]interface{}{"email": email}).Error

	if errors.Is(err,gorm.ErrRecordNotFound){
		return nil, err
	}

	return &user, nil
}

func NewUserRepository() UserRepository {
	return &userRepository{database: util.DB}
}
