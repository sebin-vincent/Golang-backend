package controller

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/dto"
	"github.com/wallet-tracky/Golang-backend/dto/response"
	"github.com/wallet-tracky/Golang-backend/service"
	"strings"
)

type CategoryController struct {
	service service.CategoryService
}

func (controller *CategoryController) GetCategories(ctx *gin.Context) {
	var categories *[]response.Category

	categoryType := ctx.Param("type")
	categoryType = strings.ToUpper(categoryType)

	if categoryType == "EXPENSE" || categoryType == "INCOME" {
		logger.Debugf("Request to fetch %S categories", categoryType)
		categories = controller.service.GetCategories(categoryType)
	} else {
		logger.Info("Invalid value for category type :", categoryType)
		ctx.AbortWithStatusJSON(400, dto.ErrorResponse{Status: 400, Message: "Invalid value for category type"})
	}

	ctx.JSON(200, categories)

}

func NewCategoryController(service service.CategoryService) CategoryController {

	return CategoryController{service: service}
}
