package usecases

import (
	"context"

	"github.com/lawmatsuyama/pismo-transactions/domain"
)

type TransactionUseCase struct {
	transaction domain.TransactionRepository
	account     domain.AccountRepository
}

func NewTransactionUseCase(transactionRepository domain.TransactionRepository, accountRepository domain.AccountRepository) TransactionUseCase {
	return TransactionUseCase{
		transaction: transactionRepository,
		account:     accountRepository,
	}
}

func (useCase TransactionUseCase) Create(ctx context.Context, tr domain.Transaction) (id string, err error) {
	if err = tr.IsValid(); err != nil {
		return
	}

	_, err = useCase.account.Get(ctx, domain.AccountFilter{ID: tr.AccountID})
	if err != nil {
		return
	}

	tr.SetID()
	tr.SetAmountSign()
	tr.SetEventDate()

	err = useCase.transaction.Create(ctx, tr)
	if err != nil {
		return "", err
	}

	return tr.ID, nil
}
