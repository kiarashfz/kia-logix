package mappers

import (
	"kia-logix/internal/user"
)

func UserArgsMapToDomain(userMap map[string]interface{}) *user.User {
	return &user.User{
		ID:      uint(userMap["ID"].(float64)),
		Name:    userMap["Name"].(string),
		Phone:   userMap["Phone"].(string),
		IsAdmin: userMap["IsAdmin"].(bool),
	}
}
