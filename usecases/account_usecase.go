package usecases

import (
	"context"

	"github.com/lawmatsuyama/pismo-transactions/domain"
)

type AccountUseCase struct {
	account domain.AccountRepository
}

func NewAccountUseCase(repository domain.AccountRepository) AccountUseCase {
	return AccountUseCase{account: repository}
}

func (useCase AccountUseCase) Create(ctx context.Context, acc domain.Account) (id string, err error) {
	if err = acc.IsValid(); err != nil {
		return
	}

	acc.SetID()
	acc.SetCreatedAt()
	err = useCase.account.Create(ctx, acc)
	if err != nil {
		return
	}

	return acc.ID, nil
}

func (useCase AccountUseCase) Get(ctx context.Context, filter domain.AccountFilter) (accs []domain.Account, err error) {
	if err = filter.IsValid(); err != nil {
		return
	}

	return useCase.account.Get(ctx, filter)
}
