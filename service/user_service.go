package service

import (
	"errors"
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/dto/response"
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/repository"
	"github.com/wallet-tracky/Golang-backend/util"
	"time"
)

type UserService interface {
	UserSignup(userSignupRequest *request.UserSignUp) (string,error)
	Login(loginRequest request.LoginRequest) (*response.LoginResponse,error)
}

type userService struct {
	repo repository.UserRepository
}

func (service *userService) Login(loginRequest request.LoginRequest) (*response.LoginResponse,error) {

	loginResponse := &response.LoginResponse{}

	user, err := service.repo.FindByEmail(loginRequest.Email)

	if err!=nil{
		return nil,err
	}

	isPasswordMatching:=user.ComparePassword(loginRequest.Password)

	if !isPasswordMatching{
		return nil,errors.New("invalid email/password")
	}

	accessToken, err := util.GenerateToken(user.Id, "AccessToken")
	if err!=nil{
		logger.Error("Access Token generation failed")
		return nil,err
	}

	refreshToken, err := util.GenerateToken(user.Id, "AccessToken")
	if err!=nil{
		logger.Error("Refresh Token generation failed")
		return nil,err
	}

	loginResponse.AccessToken=accessToken
	loginResponse.RefreshToken=refreshToken


	return loginResponse,nil
}

func (service *userService) UserSignup(userSignUpRequest *request.UserSignUp) (string,error){

	userModel := makeUserModelFromDTO(userSignUpRequest)

	encryptedPassword,err:=util.EncryptPassword(userModel.Password)

	if err!=nil{
		logger.Error("Error while hashing password")
		logger.Error(err.Error())
		return "",errors.New("something went wrong")
	}

	userModel.Password=encryptedPassword

	err = service.repo.Save(userModel)

	if err!=nil {
		logger.Error(err.Error())
		return "SomeThing wrong",err
	}

	return "User saved",nil
}

func makeUserModelFromDTO(signUpRequest *request.UserSignUp) *model.User {

	now := time.Now().Format(util.TIME_LAYOUT)

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
