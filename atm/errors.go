package atm

// список ошибок

import "errors"

var (
	ErrInvalidAmount 	 = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrAccountNotFound   = errors.New("account not found")
	ErrInvalidAccountID  = errors.New("invalid account id")
)
