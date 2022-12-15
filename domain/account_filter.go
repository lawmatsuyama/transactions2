package domain

import "time"

type AccountFilter struct {
	ID             string         `json:"id"`
	DocumentNumber DocumentNumber `json:"document_number"`
	CreatedAtFrom  time.Time      `json:"created_at_from"`
	CreatedAtTo    time.Time      `json:"created_at_to"`
	Paging         *Paging        `json:"paging" bson:"paging"`
}

func (filter AccountFilter) IsValid() error {
	if filter.DocumentNumber != "" {
		return filter.DocumentNumber.IsValid()
	}
	return nil
}
