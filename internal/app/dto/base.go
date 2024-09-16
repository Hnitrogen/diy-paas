package dto

type UserLogin struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserCreateInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserToken struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"jwt_token"`
}
