package apimanager

import "github.com/lawmatsuyama/pismo-transactions/domain"

// GenericResponse represents a generic response to be used by all api operations
type GenericResponse[T any] struct {
	Error  string `json:"error,omitempty"`
	Result T      `json:"result"`
}

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

func (request CreateAccountRequest) ToAccount() domain.Account {
	return domain.NewAccount(domain.DocumentNumber(request.DocumentNumber))
}

type CreateAccountResponse struct {
	AccountID string `json:"account_id"`
}

func FromAccountID(id string) CreateAccountResponse {
	return CreateAccountResponse{AccountID: id}
}

type GetAccountResponse struct {
	AccountID      string `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func FromAccount(acc domain.Account) GetAccountResponse {
	return GetAccountResponse{
		AccountID:      acc.ID,
		DocumentNumber: acc.DocumentNumber.String(),
	}
}
