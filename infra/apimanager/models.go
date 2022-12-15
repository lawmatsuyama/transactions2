package apimanager

import (
	"time"

	"github.com/lawmatsuyama/pismo-transactions/domain"
)

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

type CreateTransactionRequest struct {
	AccountID       string  `json:"account_id"`
	Description     string  `json:"description"`
	OperationTypeID int64   `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func (request CreateTransactionRequest) ToTransaction() domain.Transaction {
	return domain.Transaction{
		AccountID:       request.AccountID,
		Description:     request.Description,
		OperationTypeID: domain.OperationType(request.OperationTypeID),
		Amount:          request.Amount,
	}
}

type CreateTransactionResponse struct {
	ID string `json:"id"`
}

func FromTransaction(tr domain.Transaction) CreateTransactionResponse {
	return CreateTransactionResponse{
		ID: tr.ID,
	}
}

// Paging represents paging struct
type Paging struct {
	Page     int64  `json:"page"`
	NextPage *int64 `json:"next_page,omitempty"`
}

type GetTransactionRequest struct {
	ID              string    `json:"id"`
	AccountID       string    `json:"account_id"`
	Description     string    `json:"description"`
	OperationTypeID int64     `json:"operation_type_id"`
	AmountGreater   float64   `json:"amount_greater"`
	AmountLess      float64   `json:"amount_less"`
	EventDateFrom   time.Time `json:"event_date_from"`
	EventDateTo     time.Time `json:"event_date_to"`
	Paging          *Paging   `json:"paging,omitempty"`
}

func (request GetTransactionRequest) ToTransaction() domain.TransactionFilter {
	filter := domain.TransactionFilter{
		ID:              request.ID,
		AccountID:       request.AccountID,
		Description:     request.Description,
		OperationTypeID: domain.OperationType(request.OperationTypeID),
		AmountGreater:   request.AmountGreater,
		AmountLess:      request.AmountLess,
		EventDateFrom:   request.EventDateFrom,
		EventDateTo:     request.EventDateTo,
	}

	if request.Paging != nil {
		filter.Paging = &domain.Paging{
			Page: request.Paging.Page,
		}
	}
	return filter
}

// Transaction represents a transaction of account
type Transaction struct {
	ID              string  `json:"id"`
	AccountID       string  `json:"account_id"`
	Description     string  `json:"description"`
	Amount          float64 `json:"amount"`
	OperationTypeID string  `json:"operation_type_id"`
	EventDate       string  `json:"event_date"`
}

type GetTransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
	Paging       *Paging       `json:"paging" bson:"paging"`
}

func FromTransactionPaging(trsPag domain.TransactionsPaging) GetTransactionResponse {
	trs := make([]Transaction, len(trsPag.Transactions))
	for i, tr := range trsPag.Transactions {
		trs[i] = Transaction{
			ID:              tr.ID,
			AccountID:       tr.AccountID,
			Description:     tr.Description,
			OperationTypeID: tr.OperationTypeID.String(),
			Amount:          tr.Amount,
			EventDate:       domain.TimeSaoPaulo(tr.EventDate).Format(time.RFC3339),
		}
	}

	trsGetResp := GetTransactionResponse{Transactions: trs}
	if trsPag.Paging != nil {
		trsGetResp.Paging = &Paging{
			Page:     trsPag.Paging.Page,
			NextPage: trsPag.Paging.NextPage,
		}
	}

	return trsGetResp
}
