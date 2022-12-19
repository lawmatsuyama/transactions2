package usecases

import (
	"context"

	"github.com/lawmatsuyama/pismo-transactions/domain"
	log "github.com/sirupsen/logrus"
)

// TransactionUseCase implements domain.TransactionUseCase interface
type TransactionUseCase struct {
	transaction domain.TransactionRepository
	account     domain.AccountRepository
}

// NewTransactionUseCase returns a new TransactionUseCase
func NewTransactionUseCase(transactionRepository domain.TransactionRepository, accountRepository domain.AccountRepository) TransactionUseCase {
	return TransactionUseCase{
		transaction: transactionRepository,
		account:     accountRepository,
	}
}

// Create check if transaction is valid and create it in application
func (useCase TransactionUseCase) Create(ctx context.Context, tr domain.Transaction) (id string, err error) {
	l := log.WithField("transaction", tr)
	if err = tr.IsValid(); err != nil {
		l.WithError(err).Info("transaction is invalid")
		return
	}

	accs, err := useCase.account.Get(ctx, domain.AccountFilter{ID: tr.AccountID})
	if err != nil {
		l.WithError(err).Error("failed to get account")
		return
	}

	if len(accs) == 0 {
		l.Info("not found account")
		return "", domain.ErrAccountNotFound
	}

	acc := accs[0]

	tr.SetID()
	tr.SetAmountSign()
	tr.SetCurrentTimeToEventDate()
	if !acc.HasLimit(tr) {
		return "", domain.ErrHasNoLimit
	}

	acc.SetLimitByTransaction(tr)

	err = useCase.transaction.Create(ctx, tr)
	if err != nil {
		l.WithError(err).Error("failed to create transaction")
		return "", err
	}

	err = useCase.account.Update(ctx, acc)
	if err != nil {
		l.WithError(err).Error("failed to update account")
		return "", err
	}

	return tr.ID, nil
}

// Get returns transactions by given TransactionFilter
func (useCase TransactionUseCase) Get(ctx context.Context, filter domain.TransactionFilter) (trsPag domain.TransactionsPaging, err error) {
	l := log.WithField("filter", filter)
	if err = filter.IsValid(); err != nil {
		l.WithError(err).Info("filter is invalid")
		return
	}

	trs, err := useCase.transaction.Get(ctx, filter)
	if err != nil {
		l.WithError(err).Error("failed to get transaction")
		return
	}

	trsPag = domain.NewTransactionsPage(trs, filter.Paging)
	if err = trsPag.IsValid(); err != nil {
		l.WithError(err).Error("transaction page is invalid")
		return
	}

	trsPag.SetNextPaging()

	return
}
