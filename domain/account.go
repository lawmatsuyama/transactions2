package domain

import "time"

type Account struct {
	ID             string         `json:"account_id" bson:"_id"`
	DocumentNumber DocumentNumber `json:"document_number" bson:"document_number"`
	CreatedAt      time.Time      `json:"created_at" bson:"created_at"`
}

func NewAccount(docNum DocumentNumber) Account {
	return Account{DocumentNumber: docNum}
}

func (acc Account) IsValid() error {
	if err := acc.DocumentNumber.IsValid(); err != nil {
		return err
	}

	return nil
}

func (acc *Account) SetID() {
	if acc.ID == "" {
		acc.ID = UUID.Generate()
	}
}

func (acc *Account) SetCreatedAt() {
	acc.CreatedAt = Now()
}
