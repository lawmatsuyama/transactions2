package domain

import "time"

// TransactionFilter represents a filter to query transactions
type TransactionFilter struct {
	ID              string        `json:"id"`
	AccountID       string        `json:"account_id"`
	Description     string        `json:"description"`
	OperationTypeID OperationType `json:"operation_type_id"`
	AmountGreater   float64       `json:"amount_greater"`
	AmountLess      float64       `json:"amount_less"`
	EventDateFrom   time.Time     `json:"event_date_from"`
	EventDateTo     time.Time     `json:"event_date_to"`
	Paging          *Paging       `json:"paging,omitempty"`
}

// Validate valids transaction filter
func (tr TransactionFilter) IsValid() error {
	if tr.OperationTypeID != 0 {
		if err := tr.OperationTypeID.IsValid(); err != nil {
			return err
		}
	}

	return nil
}
