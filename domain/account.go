package domain

import "time"

// Account represents an user account
type Account struct {
	ID             string         `json:"account_id" bson:"_id"`
	DocumentNumber DocumentNumber `json:"document_number" bson:"document_number"`
	CreatedAt      time.Time      `json:"created_at" bson:"created_at"`
}

// NewAccount returns a new Account
func NewAccount(docNum DocumentNumber) Account {
	return Account{DocumentNumber: docNum}
}

// IsValid returns an error when account is invalid
func (acc Account) IsValid() error {
	if err := acc.DocumentNumber.IsValid(); err != nil {
		return err
	}

	return nil
}

// SetID set uuid in account ID
func (acc *Account) SetID() {
	if acc.ID == "" {
		acc.ID = UUID.Generate()
	}
}

// SetCurrentTimeToCreatedAt set current time in account created at
func (acc *Account) SetCurrentTimeToCreatedAt() {
	acc.CreatedAt = Now()
}
