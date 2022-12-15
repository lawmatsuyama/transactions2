package apimanager_test

import (
	"context"
	"net/http"

	"github.com/lawmatsuyama/pismo-transactions/domain"
)

var (
	WriteMock       func([]byte) (int, error)
	WriteHeaderMock func(statusCode int)
)

type mockResponseWriter struct{}

func (m mockResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (m mockResponseWriter) Write(b []byte) (int, error) {
	return WriteMock(b)
}
func (m mockResponseWriter) WriteHeader(statusCode int) {
	WriteHeaderMock(statusCode)
}

var (
	CreateAccountMock func(ctx context.Context, acc domain.Account) (id string, err error)
	GetAccountMock    func(ctx context.Context, filter domain.AccountFilter) (accs []domain.Account, err error)
)

type mockAccount struct{}

func (m mockAccount) Create(ctx context.Context, acc domain.Account) (id string, err error) {
	return CreateAccountMock(ctx, acc)
}

func (m mockAccount) Get(ctx context.Context, filter domain.AccountFilter) (accs []domain.Account, err error) {
	return GetAccountMock(ctx, filter)
}
