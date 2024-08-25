package handlers

import (
	"context"
)

type ServiceFactory[T any] func(context.Context) T

//func BodyValidator[T any](req T) error {
//	myValidator := presenter.GetValidator()
//	if errs := myValidator.Validate(req); len(errs) > 0 {
//		errMsgs := make([]string, 0)
//
//		for _, err := range errs {
//			errMsgs = append(errMsgs, err.Error)
//		}
//
//		return errors.New(strings.Join(errMsgs, "and"))
//	}
//	return nil
//}
