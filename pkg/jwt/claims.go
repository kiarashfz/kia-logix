package jwt

import (
	jwtLib "github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	jwtLib.RegisteredClaims
	UserID uint
}
