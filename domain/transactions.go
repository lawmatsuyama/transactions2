package domain

import (
	"math"
	"time"
)

type Transaction struct {
	ID          string        `json:"id" bson:"id"`
	AccountID   string        `json:"account_id" bson:"account_id"`
	OperationID OperationType `json:"operation_id" bson:"operation_id"`
	Amount      float64       `json:"amount" bson:"amount"`
	EventDate   time.Time     `json:"event_date" bson:"event_date"`
}

func (tr Transaction) IsValid() error {
	if tr.Amount == 0 {
		return ErrTransactionZeroAmount
	}

	return tr.OperationID.IsValid()
}

func (tr *Transaction) SetAmountSign() {
	tr.Amount = math.Abs(tr.Amount) * tr.OperationID.Sign()
}
