package domain

import (
	"math"
	"time"
)

// Transaction represents a transaction of an account
type Transaction struct {
	ID              string        `json:"id" bson:"_id"`
	AccountID       string        `json:"account_id" bson:"account_id"`
	Description     string        `json:"description" bson:"description"`
	OperationTypeID OperationType `json:"operation_type_id" bson:"operation_type_id"`
	Amount          float64       `json:"amount" bson:"amount"`
	EventDate       time.Time     `json:"event_date" bson:"event_date"`
}

// IsValid check if transaction is valid
func (tr Transaction) IsValid() error {
	if tr.Amount == 0 {
		return ErrTransactionZeroAmount
	}

	return tr.OperationTypeID.IsValid()
}

// SetAmountSign set amount sign according the operation type
func (tr *Transaction) SetAmountSign() {
	tr.Amount = math.Abs(tr.Amount) * tr.OperationTypeID.Sign()
}

// SetID generate and set a new UUID in transaction ID
func (tr *Transaction) SetID() {
	if tr.ID == "" {
		tr.ID = UUID.Generate()
	}
}

// SetCurrentTimeToEventDate set the current time in transaction event date
func (tr *Transaction) SetCurrentTimeToEventDate() {
	tr.EventDate = Now()
}
