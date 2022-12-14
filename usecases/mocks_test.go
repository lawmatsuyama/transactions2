package usecases_test

import (
	"context"

	"github.com/lawmatsuyama/pismo-transactions/domain"
)

var (
	CreateAccountMock func(ctx context.Context, acc domain.Account) (err error)
	GetAccountMock    func(ctx context.Context, filter domain.AccountFilter) (accs []domain.Account, err error)
	UUIDGenerateMock  func() string
)

type mockAccount struct{}

func (m mockAccount) Create(ctx context.Context, acc domain.Account) (err error) {
	return CreateAccountMock(ctx, acc)
}

func (m mockAccount) Get(ctx context.Context, filter domain.AccountFilter) (accs []domain.Account, err error) {
	return GetAccountMock(ctx, filter)
}

type mockUUID struct{}

func (m mockUUID) Generate() string {
	return UUIDGenerateMock()
}
