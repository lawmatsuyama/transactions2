package domain

import "time"

// Account represents an user account
type Account struct {
	ID                   string         `json:"account_id" bson:"_id"`
	DocumentNumber       DocumentNumber `json:"document_number" bson:"document_number"`
	AvailableCreditLimit float64        `json:"available_credit_limit" bson:"available_credit_limit"`
	CreatedAt            time.Time      `json:"created_at" bson:"created_at"`
}

// NewAccount returns a new Account
func NewAccount(docNum DocumentNumber, availableCreditLimit float64) Account {
	return Account{DocumentNumber: docNum, AvailableCreditLimit: availableCreditLimit}
}

// IsValid returns an error when account is invalid
func (acc Account) IsValid() error {
	if err := acc.DocumentNumber.IsValid(); err != nil {
		return err
	}

	if acc.AvailableCreditLimit < 0 {
		return ErrLimitIsNegative
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

func (acc Account) HasLimit(tr Transaction) bool {
	return acc.AvailableCreditLimit+tr.Amount >= 0
}

func (acc *Account) SetLimitByTransaction(tr Transaction) {
	acc.AvailableCreditLimit += tr.Amount
}
