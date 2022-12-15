package domain

import (
	"math"
	"time"
)

type Transaction struct {
	ID              string        `json:"id" bson:"_id"`
	AccountID       string        `json:"account_id" bson:"account_id"`
	Description     string        `json:"description" bson:"description"`
	OperationTypeID OperationType `json:"operation_type_id" bson:"operation_type_id"`
	Amount          float64       `json:"amount" bson:"amount"`
	EventDate       time.Time     `json:"event_date" bson:"event_date"`
}

func (tr Transaction) IsValid() error {
	if tr.Amount == 0 {
		return ErrTransactionZeroAmount
	}

	return tr.OperationTypeID.IsValid()
}

func (tr *Transaction) SetAmountSign() {
	tr.Amount = math.Abs(tr.Amount) * tr.OperationTypeID.Sign()
}

func (tr *Transaction) SetID() {
	if tr.ID == "" {
		tr.ID = UUID.Generate()
	}
}

func (tr *Transaction) SetEventDate() {
	tr.EventDate = Now()
}
