package helpers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate = validator.New()

type BodyValidator struct {
	validator *validator.Validate
}

func (v BodyValidator) Validate(data interface{}) []error {
	var validationErrors []error

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err)
		}
	}

	return validationErrors
}

var appValidator *BodyValidator

func GetValidator() *BodyValidator {
	if appValidator == nil {
		appValidator = &BodyValidator{
			validator: validate,
		}
	}
	return appValidator
}

func ValidateBody[T any](req T) error {
	myValidator := GetValidator()
	if errs := myValidator.Validate(req); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, err.Error())
		}

		return errors.New(strings.Join(errMsgs, "and"))
	}
	return nil
}
