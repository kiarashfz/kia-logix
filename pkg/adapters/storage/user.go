package storage

import (
	"context"
	"errors"
	"kia-logix/internal/user"
	"kia-logix/pkg/adapters/storage/entities"
	"kia-logix/pkg/adapters/storage/mappers"
	"kia-logix/pkg/errors"
	"strings"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) user.Repo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, user *user.User) (*user.User, error) {
	newUser := mappers.UserDomainToEntity(user)
	err := r.db.WithContext(ctx).Create(&newUser).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, internalErrors.DBErrDuplicateKey
		}
		return nil, err
	}
	createdUser := mappers.UserEntityToDomain(newUser)
	return createdUser, nil
}

func (r *userRepo) GetByID(ctx context.Context, id uint) (*user.User, error) {
	var userEntity entities.User

	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("id = ?", id).First(&userEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&userEntity), nil
}

func (r *userRepo) GetByPhone(ctx context.Context, phone string) (*user.User, error) {
	var userEntity entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("phone = ?", phone).First(&userEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&userEntity), nil
}
