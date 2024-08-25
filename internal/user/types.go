package user

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Repo interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, id uint) (*User, error)
	GetByPhone(ctx context.Context, phone string) (*User, error)
}

type User struct {
	ID       uint
	Name     string
	Phone    string
	Password string
	IsAdmin  bool
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) IsValidPassword(password string) bool {
	h := sha256.New()
	h.Write([]byte(password))
	passSha256 := h.Sum(nil)
	return fmt.Sprintf("%x", passSha256) == u.Password
}

func ValidatePhone(phone string) error {
	pattern := `^((0?9)|(\+?989))\d{2}\W?\d{3}\W?\d{4}$`

	re := regexp.MustCompile(pattern)

	if !re.MatchString(phone) {
		return errors.New("invalid phone number format")
	}

	return nil
}

func NormalizePhone(phone string) string {
	// Remove all non-digit characters
	re := regexp.MustCompile(`\D`)
	normalizedPhone := re.ReplaceAllString(phone, "")

	if strings.HasPrefix(normalizedPhone, "989") {
		normalizedPhone = "0" + normalizedPhone[2:]
	}
	if len(normalizedPhone) == 10 {
		normalizedPhone = "0" + normalizedPhone
	}
	if len(normalizedPhone) == 10 {
		normalizedPhone = "0" + normalizedPhone
	}
	if normalizedPhone[:4] == "0098" {
		normalizedPhone = "0" + normalizedPhone[4:]
	}

	return normalizedPhone
}

func ValidatePassword(password string) error {
	// Check for at least one digit
	if match, _ := regexp.MatchString(`[0-9]`, password); !match {
		return ErrInvalidPassword
	}

	// Check for at least one uppercase letter
	if match, _ := regexp.MatchString(`[A-Z]`, password); !match {
		return ErrInvalidPassword
	}

	// Check for at least one lowercase letter
	if match, _ := regexp.MatchString(`[a-z]`, password); !match {
		return ErrInvalidPassword
	}

	// Check for at least one special character
	if match, _ := regexp.MatchString(`[^\w\d\s:]`, password); !match {
		return ErrInvalidPassword
	}

	// Check password length
	if len(password) < 8 || len(password) > 16 {
		return ErrInvalidPassword
	}

	return nil
}
