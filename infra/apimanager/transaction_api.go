package apimanager

import (
	"context"
	"net/http"

	"github.com/lawmatsuyama/pismo-transactions/domain"
)

// TransactionAPI represents an API for transaction
type TransactionAPI struct {
	Transaction domain.TransactionUseCase
}

func NewTransactionAPI(tr domain.TransactionUseCase) TransactionAPI {
	return TransactionAPI{Transaction: tr}
}

func (api TransactionAPI) Create(w http.ResponseWriter, r *http.Request) {
	request, err := Decode[CreateTransactionRequest](r.Body)
	if err != nil {
		HandleResponse[*string](w, r, nil, domain.ErrInvalidAccount)
		return
	}

	ctx := context.Background()
	id, err := api.Transaction.Create(ctx, request.ToTransaction())
	if err != nil {
		HandleResponse[*string](w, r, nil, err)
		return
	}

	HandleResponse(w, r, FromAccountID(id), err)
}

func (api TransactionAPI) Get(w http.ResponseWriter, r *http.Request) {
	request, err := Decode[GetTransactionRequest](r.Body)
	if err != nil {
		HandleResponse[*string](w, r, nil, domain.ErrInvalidTransaction)
		return
	}

	ctx := context.Background()
	trsPag, err := api.Transaction.Get(ctx, request.ToTransaction())
	if err != nil {
		HandleResponse[*string](w, r, nil, err)
		return
	}

	HandleResponse(w, r, FromTransactionPaging(trsPag), err)
}
