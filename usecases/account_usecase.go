package usecases

import (
	"context"

	"github.com/lawmatsuyama/pismo-transactions/domain"
	log "github.com/sirupsen/logrus"
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
	l := log.WithField("account", acc)
	if err = acc.IsValid(); err != nil {
		l.WithError(err).Info("account is invalid")
		return
	}

	acc.SetID()
	acc.SetCurrentTimeToCreatedAt()
	err = useCase.account.Create(ctx, acc)
	if err != nil {
		l.WithError(err).Error("failed to create account")
		return
	}

	return acc.ID, nil
}

// Get returns accounts by given AccountFilter
func (useCase AccountUseCase) Get(ctx context.Context, filter domain.AccountFilter) (accs []domain.Account, err error) {
	l := log.WithField("filter", filter)
	if err = filter.IsValid(); err != nil {
		l.WithError(err).Info("filter is invalid")
		return
	}

	accs, err = useCase.account.Get(ctx, filter)
	if err != nil {
		l.WithError(err).Error("failed to get account")
	}
	return
}
