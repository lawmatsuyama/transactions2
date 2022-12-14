package domain

import "time"

type Account struct {
	ID        DocumentNumber `json:"id" bson:"_id"`
	CreatedAt time.Time      `json:"created_at" bson:"created_at"`
}

func (acc Account) IsValid() error {
	if err := acc.ID.IsValid(); err != nil {
		return err
	}

	return nil
}
