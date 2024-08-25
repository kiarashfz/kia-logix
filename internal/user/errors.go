package user

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrPhoneAlreadyExists    = errors.New("phone already exists")
	ErrInvalidPhone          = errors.New("phone format is not valid")
	ErrInvalidPassword       = errors.New("password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character")
	ErrInvalidAuthentication = errors.New("phone and password doesn't match")
)
