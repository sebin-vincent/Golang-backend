package service

import (
	"github.com/wallet-tracky/Golang-backend/dto/response"
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/repository"
)

type CategoryService interface {
	GetCategories(categoryType string) *[]response.Category
}

type categoryService struct {
	repo repository.CategoryRepository
}

func (service *categoryService) GetCategories(categoryType string) *[]response.Category {
	categories := service.repo.FindByType(categoryType)

	categoryResponse:=make([]response.Category,len(*categories))

	for index, category := range *categories {
		categoryDTO:= makeNewCategoryResponseDTO(&category)
		categoryResponse[index] = *categoryDTO
	}

	return &categoryResponse
}

func makeNewCategoryResponseDTO(category *model.Category) *response.Category {

	categoryResponse := new(response.Category)

	categoryResponse.Id = category.Id
	categoryResponse.Name =category.Name

	return categoryResponse
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService{
	return &categoryService{repo: repo}
}
