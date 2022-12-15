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

	return ErrUnknown
}

var (
	// ErrInvalidDocumentNumber is an error to return when there is an invalid document
	ErrInvalidDocumentNumber = ErrorTransaction{ErrorOrigin: errors.New("invalid document number"), StatusCode: http.StatusBadRequest}
	// ErrInvalidAccount is an error to return when there is an invalid account
	ErrInvalidAccount = ErrorTransaction{ErrorOrigin: errors.New("invalid account"), StatusCode: http.StatusBadRequest}
	// ErrTransactionZeroAmount is an error to return when there is a transaction with zero amount
	ErrTransactionZeroAmount = ErrorTransaction{ErrorOrigin: errors.New("transaction amount is zero"), StatusCode: http.StatusBadRequest}
	// ErrInvalidOperationType is an error to return when there is a transaction with invalid operation type
	ErrInvalidOperationType = ErrorTransaction{ErrorOrigin: errors.New("invalid transaction operation type"), StatusCode: http.StatusBadRequest}
	// ErrInvalidTransaction is an error to return when there is a invalid transaction
	ErrInvalidTransaction = ErrorTransaction{ErrorOrigin: errors.New("invalid transaction"), StatusCode: http.StatusBadRequest}
	// ErrTransactionsNotFound is an error to return when no transactions found
	ErrTransactionsNotFound = ErrorTransaction{ErrorOrigin: errors.New("transactions not found"), StatusCode: http.StatusNotFound}
	// ErrAccountNotFound is an error to return when no accounts found
	ErrAccountNotFound = ErrorTransaction{ErrorOrigin: errors.New("account not found"), StatusCode: http.StatusNotFound}
	// ErrUnknown is an error to return when service get an unknown error
	ErrUnknown = ErrorTransaction{ErrorOrigin: errors.New("unknow error"), StatusCode: http.StatusBadRequest}
)
