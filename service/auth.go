package service

import (
	"context"
	"kia-logix/config"
	"kia-logix/internal/user"
	"kia-logix/pkg/jwt"
	"time"

	jwtLib "github.com/golang-jwt/jwt/v5"
)

type IAuthService interface {
	CreateUser(ctx context.Context, user *user.User) (*user.User, error)
	Login(ctx context.Context, email, pass string) (*UserToken, error)
	RefreshAuth(ctx context.Context, refreshToken string) (*UserToken, error)
}

type AuthService struct {
	userOps    user.IUserOps
	authConfig config.Auth
}

func NewAuthService(userOps *user.Ops, authConfig config.Auth) *AuthService {
	return &AuthService{
		userOps:    userOps,
		authConfig: authConfig,
	}
}

type UserToken struct {
	AuthorizationToken string `json:"auth_token"`
	RefreshToken       string `json:"refresh_token"`
}

func (s *AuthService) CreateUser(ctx context.Context, user *user.User) (*user.User, error) {
	return s.userOps.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, email, pass string) (*UserToken, error) {
	fetchedUser, err := s.userOps.GetUserByPhoneAndPassword(ctx, email, pass)
	if err != nil {
		return nil, err
	}

	// calc expiration time values
	var (
		authExp    = time.Now().Add(time.Minute * time.Duration(s.authConfig.TokenExpMinutes))
		refreshExp = time.Now().Add(time.Minute * time.Duration(s.authConfig.RefreshTokenExpMinutes))
	)

	authToken, err := jwt.CreateToken([]byte(s.authConfig.SecretToken), s.userClaims(fetchedUser, authExp))
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.CreateToken([]byte(s.authConfig.SecretToken), s.userClaims(fetchedUser, refreshExp))
	if err != nil {
		return nil, err
	}

	return &UserToken{
		AuthorizationToken: authToken,
		RefreshToken:       refreshToken,
	}, nil
}

func (s *AuthService) RefreshAuth(ctx context.Context, refreshToken string) (*UserToken, error) {
	claim, err := jwt.ParseToken(refreshToken, []byte(s.authConfig.SecretToken))
	if err != nil {
		return nil, err
	}

	u, err := s.userOps.GetUserByID(ctx, claim.UserID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, user.ErrUserNotFound
	}

	// calc expiration time values
	var (
		authExp = time.Now().Add(time.Minute * time.Duration(s.authConfig.TokenExpMinutes))
	)

	authToken, err := jwt.CreateToken([]byte(s.authConfig.SecretToken), s.userClaims(u, authExp))
	if err != nil {
		return nil, err
	}

	return &UserToken{
		AuthorizationToken: authToken,
		RefreshToken:       refreshToken,
	}, nil
}

func (s *AuthService) userClaims(user *user.User, exp time.Time) *jwt.UserClaims {
	return &jwt.UserClaims{
		RegisteredClaims: jwtLib.RegisteredClaims{
			ExpiresAt: &jwtLib.NumericDate{
				Time: exp,
			},
		},
		UserID: user.ID,
	}
}
