package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/lawmatsuyama/pismo-transactions/domain"
	"github.com/lawmatsuyama/pismo-transactions/usecases"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	testCases := []struct {
		Name                                    string
		FakeTransactionFile                     string
		FakeTransactionID                       string
		FakeErrCreateTransactionRepository      error
		FakeErrGetAccountRepository             error
		ExpInputGetAccountRepositoryFile        string
		ExpInputCreateTransactionRepositoryFile string
		ExpTransactionID                        string
		ExpErr                                  error
	}{
		{
			Name:                                    "01_should_create_transaction_return_nil_error",
			FakeTransactionFile:                     "./testdata/transaction/create/01_should_create_transaction_return_nil_error/fake_transaction.json",
			FakeTransactionID:                       "123456",
			ExpInputGetAccountRepositoryFile:        "./testdata/transaction/create/01_should_create_transaction_return_nil_error/exp_in_get_account_repo.json",
			ExpInputCreateTransactionRepositoryFile: "./testdata/transaction/create/01_should_create_transaction_return_nil_error/exp_in_create_account_repo.json",
			ExpTransactionID:                        "123456",
			ExpErr:                                  nil,
		},
		{
			Name:                                    "02_should_create_transaction_return_not_found_account",
			FakeTransactionFile:                     "./testdata/transaction/create/02_should_create_transaction_return_not_found_account/fake_transaction.json",
			FakeTransactionID:                       "123456",
			FakeErrGetAccountRepository:             domain.ErrAccountNotFound,
			ExpInputGetAccountRepositoryFile:        "./testdata/transaction/create/02_should_create_transaction_return_not_found_account/exp_in_get_account_repo.json",
			ExpInputCreateTransactionRepositoryFile: "./testdata/transaction/create/02_should_create_transaction_return_not_found_account/exp_in_create_account_repo.json",
			ExpTransactionID:                        "",
			ExpErr:                                  domain.ErrAccountNotFound,
		},
		{
			Name:                                    "03_should_create_transaction_return_unknow_error_on_create_repository",
			FakeTransactionFile:                     "./testdata/transaction/create/03_should_create_transaction_return_unknow_error_on_create_repository/fake_transaction.json",
			FakeTransactionID:                       "123456",
			FakeErrCreateTransactionRepository:      domain.ErrUnknow,
			ExpInputGetAccountRepositoryFile:        "./testdata/transaction/create/03_should_create_transaction_return_unknow_error_on_create_repository/exp_in_get_account_repo.json",
			ExpInputCreateTransactionRepositoryFile: "./testdata/transaction/create/03_should_create_transaction_return_unknow_error_on_create_repository/exp_in_create_account_repo.json",
			ExpTransactionID:                        "",
			ExpErr:                                  domain.ErrUnknow,
		},
		{
			Name:                                    "04_should_create_transaction_return_invalid_operation_type",
			FakeTransactionFile:                     "./testdata/transaction/create/04_should_create_transaction_return_invalid_operation_type/fake_transaction.json",
			FakeTransactionID:                       "123456",
			ExpInputGetAccountRepositoryFile:        "./testdata/transaction/create/04_should_create_transaction_return_invalid_operation_type/exp_in_get_account_repo.json",
			ExpInputCreateTransactionRepositoryFile: "./testdata/transaction/create/04_should_create_transaction_return_invalid_operation_type/exp_in_create_account_repo.json",
			ExpTransactionID:                        "",
			ExpErr:                                  domain.ErrInvalidOperationType,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testCreate(t, tc.Name, tc.FakeTransactionFile, tc.FakeTransactionID, tc.FakeErrCreateTransactionRepository, tc.FakeErrGetAccountRepository, tc.ExpInputGetAccountRepositoryFile, tc.ExpInputCreateTransactionRepositoryFile, tc.ExpTransactionID, tc.ExpErr)
		})
	}
}

func testCreate(t *testing.T, name, fakeTrFile, fakeTrID string, fakeErrCreateTrRepo, fakeErrGetAccRepo error, expInGetAccRepoFile, expInCreateTrRepoFile, expID string, expErr error) {
	domain.Now = func() time.Time {
		return domain.TimeSaoPaulo(time.Date(2022, 10, 10, 12, 0, 0, 0, time.UTC))
	}

	domain.UUID = mockUUID{}
	UUIDGenerateMock = func() string {
		return fakeTrID
	}

	var gotInCreateTrRepo domain.Transaction
	CreateTransactionMock = func(ctx context.Context, tr domain.Transaction) (err error) {
		gotInCreateTrRepo = tr
		return fakeErrCreateTrRepo
	}

	var gotInGetAccRepo domain.AccountFilter
	GetAccountMock = func(ctx context.Context, filter domain.AccountFilter) (accs []domain.Account, err error) {
		gotInGetAccRepo = filter
		return nil, fakeErrGetAccRepo
	}

	fakeTr := domain.ReadJSON[domain.Transaction](t, fakeTrFile)
	useCase := usecases.NewTransactionUseCase(mockTransaction{}, mockAccount{})
	gotID, gotErr := useCase.Create(context.Background(), fakeTr)
	if *update {
		domain.CreateJSON(t, expInCreateTrRepoFile, gotInCreateTrRepo)
		domain.CreateJSON(t, expInGetAccRepoFile, gotInGetAccRepo)
		return
	}

	assert.Equal(t, expErr, gotErr, "exp error should be equal got error")
	assert.Equal(t, expID, gotID, "exp ID should be equal got ID")
	domain.CompareWithFile(t, "compare input get account repository", expInGetAccRepoFile, gotInGetAccRepo)
	domain.CompareWithFile(t, "compare input create transaction repository", expInCreateTrRepoFile, gotInCreateTrRepo)

}
