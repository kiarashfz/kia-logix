package mappers

import (
	"kia-logix/internal/user"
	"kia-logix/pkg/adapters/storage/entities"
)

func UserDomainToEntity(domainUser *user.User) *entities.User {
	return &entities.User{
		Name:     domainUser.Name,
		Phone:    domainUser.Phone,
		Password: domainUser.Password,
		IsAdmin:  domainUser.IsAdmin,
	}
}

func UserEntityToDomain(userEntity *entities.User) *user.User {
	return &user.User{
		ID:       userEntity.ID,
		Name:     userEntity.Name,
		Phone:    userEntity.Phone,
		Password: userEntity.Password,
		IsAdmin:  userEntity.IsAdmin,
	}
}
