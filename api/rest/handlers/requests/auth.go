package requests

import "kia-logix/internal/user"

type UserRegister struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func UserRegisterToUserDomain(u *UserRegister) *user.User {
	return &user.User{
		Name:     u.Name,
		Phone:    u.Phone,
		Password: u.Password,
	}
}
