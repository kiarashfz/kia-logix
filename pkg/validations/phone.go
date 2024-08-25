package validations

import (
	"errors"
	"regexp"
	"strings"
)

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
