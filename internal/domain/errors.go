package domain

import "errors"

var (
	// ErrAccountNotFound é retornado quando uma conta não é encontrada
	ErrAccountNotFound = errors.New("account not found")
	// ErrDuplicated APIKey é retornado quando há tentativa de criar conta com API key duplicada
	ErrDuplicatedAPIKey = errors.New("appi key already exists")
	// ErrInvoiceNotFound é retornado quando uma fatura não é encontrada
	ErrInvoiceNotFound = errors.New("invoice not found")
	// ErrUnauthorizedAcess é retornado quando há tentativa de acesso não autorizado a um recurso
	ErrUnauthorizedAcess = errors.New("unauthorized access")

	ErrInvalidAmount = errors.New("invalid amount")
	ErrInvalidStatus = errors.New("invalid status")
)
