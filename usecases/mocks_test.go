package usecases_test

import (
	"context"

	"github.com/lawmatsuyama/pismo-transactions/domain"
)

var (
	CreateAccountMock func(ctx context.Context, acc domain.Account) (err error)
	GetAccountMock    func(ctx context.Context, filter domain.AccountFilter) (accs []domain.Account, err error)
)

type mockAccount struct{}

func (m mockAccount) Create(ctx context.Context, acc domain.Account) (err error) {
	return CreateAccountMock(ctx, acc)
}

func (m mockAccount) Get(ctx context.Context, filter domain.AccountFilter) (accs []domain.Account, err error) {
	return GetAccountMock(ctx, filter)
}

var (
	UUIDGenerateMock func() string
)

type mockUUID struct{}

func (m mockUUID) Generate() string {
	return UUIDGenerateMock()
}

var (
	CreateTransactionMock func(ctx context.Context, tr domain.Transaction) (err error)
	GetTransactionsMock   func(ctx context.Context, tr domain.TransactionFilter) ([]*domain.Transaction, error)
)

type mockTransaction struct{}

func (m mockTransaction) Create(ctx context.Context, tr domain.Transaction) (err error) {
	return CreateTransactionMock(ctx, tr)
}

func (m mockTransaction) Get(ctx context.Context, tr domain.TransactionFilter) ([]*domain.Transaction, error) {
	return GetTransactionsMock(ctx, tr)
}
