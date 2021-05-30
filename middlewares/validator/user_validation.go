package validator

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	. "github.com/wallet-tracky/Golang-backend/dto"
	"github.com/wallet-tracky/Golang-backend/dto/request"
)

type UserValidator struct {
}

func (validate *UserValidator) ValidateUserSignUp(context *gin.Context) {

	var signUpRequest request.UserSignUp

	err := context.BindJSON(&signUpRequest)

	if err != nil {
		logger.Error("Invalid request format.", err.Error())
		context.AbortWithStatusJSON(400, ErrorResponse{Status: 400, Message: err.Error()})
		return
	}

	if !isValidPassword(signUpRequest.Password) {
		logger.Error("Invalid password")
		context.AbortWithStatusJSON(400, ErrorResponse{Status: 400, Message: "Invalid password"})
		return
	}

	context.Set("request", &signUpRequest)
	context.Next()
}

func (validate *UserValidator) LoginValidator(context *gin.Context) {

	var loginRequest request.LoginRequest

	err := context.BindJSON(&loginRequest)

	if err != nil {
		logger.Error("Invalid request format.", err.Error())
		context.AbortWithStatusJSON(400, ErrorResponse{Status: 400,Message: err.Error()})
		return
	}

	context.Set("request", &loginRequest)
	context.Next()

}

func isValidPassword(password string) bool {
	isValid := true
	if len(password) < 6 {
		isValid = false
	}
	return isValid
}
