package user

import (
	"context"
	"errors"
	internalErrors "kia-logix/pkg/errors"
	"kia-logix/pkg/utils"
)

type IUserOps interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetUserByID(ctx context.Context, id uint) (*User, error)
	GetUserByPhoneAndPassword(ctx context.Context, phone, password string) (*User, error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
}

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{
		repo: repo,
	}
}

func (o *Ops) Create(ctx context.Context, user *User) (*User, error) {
	err := validateUserRegistration(user)
	if err != nil {
		return nil, err
	}

	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPass

	user.Phone = NormalizePhone(user.Phone)

	createdUser, err := o.repo.Create(ctx, user)
	if err != nil {
		if errors.Is(err, internalErrors.DBErrDuplicateKey) {
			return nil, ErrPhoneAlreadyExists
		}
		return nil, err
	}
	return createdUser, nil
}

func (o *Ops) GetUserByID(ctx context.Context, id uint) (*User, error) {
	return o.repo.GetByID(ctx, id)
}

func (o *Ops) GetUserByPhoneAndPassword(ctx context.Context, phone, password string) (*User, error) {
	phone = NormalizePhone(phone)
	user, err := o.repo.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if err = utils.CheckPasswordHash(password, user.Password); err != nil {
		return nil, ErrInvalidAuthentication
	}

	return user, nil
}

func (o *Ops) GetUserByPhone(ctx context.Context, phone string) (*User, error) {
	phone = NormalizePhone(phone)
	user, err := o.repo.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func validateUserRegistration(user *User) error {
	err := ValidatePhone(user.Phone)
	if err != nil {
		return ErrInvalidPhone
	}

	if err := ValidatePassword(user.Password); err != nil {
		return ErrInvalidPassword
	}
	return nil
}
