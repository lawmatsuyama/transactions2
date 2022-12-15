package usecases

import (
	"context"

	"github.com/lawmatsuyama/pismo-transactions/domain"
)

// AccountUseCase implements domain.AccountUseCase interface
type AccountUseCase struct {
	account domain.AccountRepository
}

// NewAccountUseCase returns a new AccountUseCase
func NewAccountUseCase(repository domain.AccountRepository) AccountUseCase {
	return AccountUseCase{account: repository}
}

// Create check if account is valid and create it in application
func (useCase AccountUseCase) Create(ctx context.Context, acc domain.Account) (id string, err error) {
	if err = acc.IsValid(); err != nil {
		return
	}

	acc.SetID()
	acc.SetCurrentTimeToCreatedAt()
	err = useCase.account.Create(ctx, acc)
	if err != nil {
		return
	}

	return acc.ID, nil
}

// Get returns accounts by given AccountFilter
func (useCase AccountUseCase) Get(ctx context.Context, filter domain.AccountFilter) (accs []domain.Account, err error) {
	if err = filter.IsValid(); err != nil {
		return
	}

	return useCase.account.Get(ctx, filter)
}
