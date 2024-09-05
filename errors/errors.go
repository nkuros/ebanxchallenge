package errors

import "errors"

var (
	ErrOriginAccountNotFound = errors.New("origin account not found")
	ErrTargetAccountNotFound = errors.New("target account not found")
	ErrInvalidAmount         = errors.New("invalid amount")
	ErrInvalidAmountFormat   = errors.New("invalid amount format")
	ErrMissingEventType      = errors.New("missing type")
	ErrMissingOriginId       = errors.New("missing origin account_id")
	ErrMissingAmount         = errors.New("missing amount")
	ErrMissingDestinationId  = errors.New("missing destination account_id")
	ErrInvalidOriginId       = errors.New("invalid origin account_id")
	ErrInvalidDestinationId  = errors.New("invalid destination account_id")
	ErrInvalidEventType      = errors.New("invalid event type")
	ErrInsufficientFunds     = errors.New("insufficient funds")
	ErrMissingBody           = errors.New("missing body")
	ErrAccountCreationFailed = errors.New("account creation failed")
)
