package domain

import "time"

// AccountFilter represents a filter to query account
type AccountFilter struct {
	ID             string         `json:"id"`
	DocumentNumber DocumentNumber `json:"document_number"`
	CreatedAtFrom  time.Time      `json:"created_at_from"`
	CreatedAtTo    time.Time      `json:"created_at_to"`
	Paging         *Paging        `json:"paging" bson:"paging"`
}

// IsValid check if account filter is valid and return error when it is invalid
func (filter AccountFilter) IsValid() error {
	if filter.DocumentNumber != "" {
		return filter.DocumentNumber.IsValid()
	}
	return nil
}
