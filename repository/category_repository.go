package repository


import (
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/util"
	"gorm.io/gorm"
)

type 	CategoryRepository interface {
	FindByType(categoryType string) *[]model.Category
}

type categoryRepository struct {
	database *gorm.DB
}

func (repo *categoryRepository) FindByType(categoryType string) *[]model.Category {

	var categories *[]model.Category

	repo.database.Find(&categories, "type=?", categoryType)

	return categories
}

func NewCategoryRepository() CategoryRepository {

	return &categoryRepository{database: util.DB}
}