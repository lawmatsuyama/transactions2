package domain

import (
	"errors"
	"net/http"
)

// ErrorTransaction represents an error and implements error interface
type ErrorTransaction struct {
	ErrorOrigin error
	StatusCode  int
}

// Error return a string error
func (err ErrorTransaction) Error() string {
	return err.ErrorOrigin.Error()
}

// Status return an http status code
func (err ErrorTransaction) Status() int {
	return err.StatusCode
}

// ErrorToErrorTransaction converts error into ErrorTransaction
func ErrorToErrorTransaction(err error) ErrorTransaction {
	errTr, ok := err.(ErrorTransaction)
	if ok {
		return errTr
	}

	return ErrUnknow
}

var (
	ErrInvalidDocumentNumber = ErrorTransaction{ErrorOrigin: errors.New("invalid document number"), StatusCode: http.StatusBadRequest}
	ErrInvalidAccount        = ErrorTransaction{ErrorOrigin: errors.New("invalid account"), StatusCode: http.StatusBadRequest}
	ErrTransactionZeroAmount = ErrorTransaction{ErrorOrigin: errors.New("transaction amount is zero"), StatusCode: http.StatusBadRequest}
	ErrInvalidOperationType  = ErrorTransaction{ErrorOrigin: errors.New("invalid transaction operation type"), StatusCode: http.StatusBadRequest}
	ErrInvalidTransaction    = ErrorTransaction{ErrorOrigin: errors.New("invalid transaction"), StatusCode: http.StatusBadRequest}
	ErrTransactionsNotFound  = ErrorTransaction{ErrorOrigin: errors.New("transactions not found"), StatusCode: http.StatusNotFound}
	ErrAccountNotFound       = ErrorTransaction{ErrorOrigin: errors.New("account not found"), StatusCode: http.StatusNotFound}
	ErrUnknow                = ErrorTransaction{ErrorOrigin: errors.New("unknow error"), StatusCode: http.StatusBadRequest}
)
