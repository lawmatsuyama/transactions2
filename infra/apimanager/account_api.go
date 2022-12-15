package apimanager

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lawmatsuyama/pismo-transactions/domain"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// AccountAPI represents an API for account
type AccountAPI struct {
	Account domain.AccountUseCase
}

// NewAccountAPI returns a new AccountAPI
func NewAccountAPI(acc domain.AccountUseCase) AccountAPI {
	return AccountAPI{Account: acc}
}

// Create godoc
//
//	@Summary		API to create account in the application.
//	@Description	Receives account data and registered it in application.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			create_account_request			body		CreateAccountRequest								true	"Create Account Request"
//	@Success		200				{object}	apimanager.GenericResponse[CreateAccountResponse]
//	@Failure		400				{object}	apimanager.GenericResponse[string]
//	@Failure		404				{object}	apimanager.GenericResponse[string]
//	@Router			/accounts [post]
func (api AccountAPI) Create(w http.ResponseWriter, r *http.Request) {
	request, err := Decode[CreateAccountRequest](r.Body)
	if err != nil {
		HandleResponse[*string](w, r, nil, domain.ErrInvalidAccount)
		return
	}
	l := log.WithField("account", request)

	ctx := context.Background()
	id, err := api.Account.Create(ctx, request.ToAccount())
	if err != nil {
		HandleResponse[*string](w, r, nil, err)
		return
	}

	HandleResponse(w, r, FromAccountID(id), err)
	l.Info("create account ok")
}

// GetByID godoc
//
//	@Summary		API to get account by ID in the application.
//	@Description	Receives path param account ID.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			accountID			path		string				true	"Account ID"
//	@Success		200				{object}	apimanager.GenericResponse[GetAccountResponse]
//	@Failure		400				{object}	apimanager.GenericResponse[string]
//	@Failure		404				{object}	apimanager.GenericResponse[string]
//	@Router			/accounts/{accountID} [get]
func (api AccountAPI) GetByID(w http.ResponseWriter, r *http.Request) {
	accID := chi.URLParam(r, "accountID")
	if accID == "" {
		HandleResponse[*string](w, r, nil, domain.ErrInvalidAccount)
		return
	}
	l := logrus.WithField("account_id", accID)
	request := domain.AccountFilter{ID: accID}
	ctx := context.Background()
	accs, err := api.Account.Get(ctx, request)
	if err != nil {
		HandleResponse[*string](w, r, nil, err)
		return
	}

	if len(accs) == 0 {
		HandleResponse[*string](w, r, nil, domain.ErrAccountNotFound)
		return
	}

	HandleResponse(w, r, FromAccount(accs[0]), err)
	l.Info("get account by id returned ok")
}
