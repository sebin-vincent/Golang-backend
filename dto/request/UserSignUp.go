package request

type UserSignUp struct {
	Email        string `json:"email" binding:"required,email"`
	Name         string `json:"name" binding:"required"`
	Password     string `json:"password"`
}


