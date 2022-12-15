package usecases

import (
	"context"

	"github.com/lawmatsuyama/pismo-transactions/domain"
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
	if err = tr.IsValid(); err != nil {
		return
	}

	accs, err := useCase.account.Get(ctx, domain.AccountFilter{ID: tr.AccountID})
	if err != nil {
		return
	}

	if len(accs) == 0 {
		return "", domain.ErrAccountNotFound
	}

	tr.SetID()
	tr.SetAmountSign()
	tr.SetCurrentTimeToEventDate()

	err = useCase.transaction.Create(ctx, tr)
	if err != nil {
		return "", err
	}

	return tr.ID, nil
}

// Get returns transactions by given TransactionFilter
func (useCase TransactionUseCase) Get(ctx context.Context, filter domain.TransactionFilter) (trsPag domain.TransactionsPaging, err error) {
	if err = filter.IsValid(); err != nil {
		return
	}

	trs, err := useCase.transaction.Get(ctx, filter)
	if err != nil {
		return
	}

	trsPag = domain.NewTransactionsPage(trs, filter.Paging)
	if err = trsPag.IsValid(); err != nil {
		return
	}

	trsPag.SetNextPaging()

	return
}
