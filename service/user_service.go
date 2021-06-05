package service

import (
	logger "github.com/sirupsen/logrus"
	. "github.com/wallet-tracky/Golang-backend/dto"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/dto/response"
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/repository"
	"github.com/wallet-tracky/Golang-backend/util"
	"time"
)

type UserService interface {
	UserSignup(userSignupRequest *request.UserSignUp) (string, *ErrorResponse)
	Login(loginRequest request.LoginRequest) (*response.LoginResponse, *ErrorResponse)
}

type userService struct {
	repo repository.UserRepository
}

func (service *userService) Login(loginRequest request.LoginRequest) (*response.LoginResponse, *ErrorResponse) {

	loginResponse := &response.LoginResponse{}

	user, err := service.repo.FindByEmail(loginRequest.Email)

	if err != nil {
		logger.Info("User does not exists with given mail")
		return nil, &ErrorResponse{Status: 404,Message: "User does not exists"}
	}

	isPasswordMatching := user.ComparePassword(loginRequest.Password)

	if !isPasswordMatching {
		return nil,&ErrorResponse{Status: 401,Message: "Email/Password does not match"}
	}

	accessToken, err := util.GenerateToken(user.Id, "AccessToken")
	if err != nil {
		logger.Error("Access Token generation failed")
		return nil, &ErrorResponse{Status: 500,Message: "Something went wrong"}
	}

	refreshToken, err := util.GenerateToken(user.Id, "AccessToken")
	if err != nil {
		logger.Error("Refresh Token generation failed")
		return nil, &ErrorResponse{Status: 500,Message: "Something went wrong"}
	}

	loginResponse.AccessToken = accessToken
	loginResponse.RefreshToken = refreshToken

	return loginResponse, nil
}

func (service *userService) UserSignup(userSignUpRequest *request.UserSignUp) (string, *ErrorResponse) {

	user, err := service.repo.FindByEmail(userSignUpRequest.Email)

	if user != nil {
		logger.Error("User with email already exists", err)
		return "", &ErrorResponse{Status: 400, Message: "User with given email already exists"}
	}

	userModel := makeUserModelFromDTO(userSignUpRequest)

	encryptedPassword, err := util.EncryptPassword(userModel.Password)

	if err != nil {
		logger.Error("Error while hashing password")
		logger.Error(err.Error())
		return "", &ErrorResponse{Status: 500, Message: "Something went wrong"}
	}

	userModel.Password = encryptedPassword

	err = service.repo.Save(userModel)

	if err != nil {
		logger.Error(err.Error())
		return "", &ErrorResponse{Status: 500, Message: "Something went wrong"}
	}

	return "User saved", nil
}

func makeUserModelFromDTO(signUpRequest *request.UserSignUp) *model.User {

	now := time.Now().Format(util.DateTimeLayout)

	userModel := &model.User{
		Email:     signUpRequest.Email,
		Name:      signUpRequest.Name,
		Password:  signUpRequest.Password,
		CreatedOn: now,
		UpdatedOn: now,
		IsEnabled: true,
		IsDeleted: false,
	}

	return userModel
}

func NewUserService(userRepository repository.UserRepository) UserService {

	service := new(userService)
	service.repo = userRepository
	return service
}
