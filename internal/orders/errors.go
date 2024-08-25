package orders

import "errors"

var (
	ErrInvalidSenderPhone   = errors.New("sender phone format is not valid")
	ErrInvalidReceiverPhone = errors.New("receiver phone format is not valid")
)
