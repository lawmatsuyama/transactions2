package apimanager

import (
	"context"
	"net/http"

	"github.com/lawmatsuyama/pismo-transactions/domain"
	log "github.com/sirupsen/logrus"
)

// TransactionAPI represents an API for transaction
type TransactionAPI struct {
	Transaction domain.TransactionUseCase
}

// NewTransactionAPI returns a new TransactionAPI
func NewTransactionAPI(tr domain.TransactionUseCase) TransactionAPI {
	return TransactionAPI{Transaction: tr}
}

// Create godoc
//
//	@Summary		API to create transaction in the application.
//	@Description	Receives transaction data and registered it in application.
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Param			create_transaction_request			body		CreateTransactionRequest								true	"Create Transaction Request"
//	@Success		200				{object}	apimanager.GenericResponse[CreateTransactionResponse]
//	@Failure		400				{object}	apimanager.GenericResponse[string]
//	@Failure		404				{object}	apimanager.GenericResponse[string]
//	@Router			/transactions [post]
func (api TransactionAPI) Create(w http.ResponseWriter, r *http.Request) {
	request, err := Decode[CreateTransactionRequest](r.Body)
	if err != nil {
		HandleResponse[*string](w, r, nil, domain.ErrInvalidAccount)
		return
	}
	l := log.WithField("transaction", request)

	ctx := context.Background()
	id, err := api.Transaction.Create(ctx, request.ToTransaction())
	if err != nil {
		HandleResponse[*string](w, r, nil, err)
		return
	}

	HandleResponse(w, r, FromTransactionID(id), err)
	l.Info("create transaction ok")
}

// Get godoc
//
//	@Summary		API to get transactions by filter.
//	@Description	List transactions by giving filter.
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Param			get_transaction_request			body		GetTransactionRequest								true	"Get Transaction Request"
//	@Success		200				{object}	apimanager.GenericResponse[GetTransactionResponse]
//	@Failure		400				{object}	apimanager.GenericResponse[string]
//	@Failure		404				{object}	apimanager.GenericResponse[string]
//	@Router			/transactions/query [post]
func (api TransactionAPI) Get(w http.ResponseWriter, r *http.Request) {
	request, err := Decode[GetTransactionRequest](r.Body)
	if err != nil {
		HandleResponse[*string](w, r, nil, domain.ErrInvalidTransaction)
		return
	}
	l := log.WithField("filter", request)
	ctx := context.Background()
	trsPag, err := api.Transaction.Get(ctx, request.ToTransaction())
	if err != nil {
		HandleResponse[*string](w, r, nil, err)
		return
	}

	HandleResponse(w, r, FromTransactionPaging(trsPag), err)
	l.Info("get transaction returned ok")
}
