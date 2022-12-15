package apimanager

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lawmatsuyama/pismo-transactions/domain"
)

// AccountAPI represents an API for account
type AccountAPI struct {
	UseCase domain.AccountUseCase
}

func NewAccountAPI(useCase domain.AccountUseCase) AccountAPI {
	return AccountAPI{UseCase: useCase}
}

func (api AccountAPI) Create(w http.ResponseWriter, r *http.Request) {
	request, err := Decode[CreateAccountRequest](r.Body)
	if err != nil {
		HandleResponse[*string](w, r, nil, domain.ErrInvalidAccount)
		return
	}

	ctx := context.Background()
	id, err := api.UseCase.Create(ctx, request.ToAccount())
	if err != nil {
		HandleResponse[*string](w, r, nil, err)
		return
	}

	HandleResponse(w, r, FromAccountID(id), err)
}

func (api AccountAPI) GetByID(w http.ResponseWriter, r *http.Request) {
	accID := chi.URLParam(r, "accountID")
	request := domain.AccountFilter{ID: accID}
	ctx := context.Background()
	accs, err := api.UseCase.Get(ctx, request)
	if err != nil {
		HandleResponse[*string](w, r, nil, err)
		return
	}

	if len(accs) == 0 {
		HandleResponse[*string](w, r, nil, domain.ErrAccountNotFound)
		return
	}

	HandleResponse(w, r, FromAccount(accs[0]), err)
}
