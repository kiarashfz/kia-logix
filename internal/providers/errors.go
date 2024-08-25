package providers

import "errors"

var (
	ErrProviderNameAlreadyExists = errors.New("provider name already exists")
	ErrProviderURLInvalid        = errors.New("invalid provider URL")
)
