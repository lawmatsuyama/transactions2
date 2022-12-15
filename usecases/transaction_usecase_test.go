package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/lawmatsuyama/pismo-transactions/domain"
	"github.com/lawmatsuyama/pismo-transactions/pkg/testhelpers"
	"github.com/lawmatsuyama/pismo-transactions/usecases"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	testCases := []struct {
		Name                                    string
		FakeTransactionFile                     string
		FakeTransactionID                       string
		FakeAccountFile                         string
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
			FakeAccountFile:                         "./testdata/transaction/create/01_should_create_transaction_return_nil_error/fake_account.json",
			ExpInputGetAccountRepositoryFile:        "./testdata/transaction/create/01_should_create_transaction_return_nil_error/exp_in_get_account_repo.json",
			ExpInputCreateTransactionRepositoryFile: "./testdata/transaction/create/01_should_create_transaction_return_nil_error/exp_in_create_transaction_repo.json",
			ExpTransactionID:                        "123456",
			ExpErr:                                  nil,
		},
		{
			Name:                                    "02_should_create_transaction_return_not_found_account",
			FakeTransactionFile:                     "./testdata/transaction/create/02_should_create_transaction_return_not_found_account/fake_transaction.json",
			FakeTransactionID:                       "123456",
			FakeAccountFile:                         "./testdata/transaction/create/02_should_create_transaction_return_not_found_account/fake_account.json",
			FakeErrGetAccountRepository:             domain.ErrAccountNotFound,
			ExpInputGetAccountRepositoryFile:        "./testdata/transaction/create/02_should_create_transaction_return_not_found_account/exp_in_get_account_repo.json",
			ExpInputCreateTransactionRepositoryFile: "./testdata/transaction/create/02_should_create_transaction_return_not_found_account/exp_in_create_transaction_repo.json",
			ExpTransactionID:                        "",
			ExpErr:                                  domain.ErrAccountNotFound,
		},
		{
			Name:                                    "03_should_create_transaction_return_unknow_error_on_create_repository",
			FakeTransactionFile:                     "./testdata/transaction/create/03_should_create_transaction_return_unknow_error_on_create_repository/fake_transaction.json",
			FakeTransactionID:                       "123456",
			FakeAccountFile:                         "./testdata/transaction/create/03_should_create_transaction_return_unknow_error_on_create_repository/fake_account.json",
			FakeErrCreateTransactionRepository:      domain.ErrUnknow,
			ExpInputGetAccountRepositoryFile:        "./testdata/transaction/create/03_should_create_transaction_return_unknow_error_on_create_repository/exp_in_get_account_repo.json",
			ExpInputCreateTransactionRepositoryFile: "./testdata/transaction/create/03_should_create_transaction_return_unknow_error_on_create_repository/exp_in_create_transaction_repo.json",
			ExpTransactionID:                        "",
			ExpErr:                                  domain.ErrUnknow,
		},
		{
			Name:                                    "04_should_create_transaction_return_invalid_operation_type",
			FakeTransactionFile:                     "./testdata/transaction/create/04_should_create_transaction_return_invalid_operation_type/fake_transaction.json",
			FakeTransactionID:                       "123456",
			FakeAccountFile:                         "./testdata/transaction/create/04_should_create_transaction_return_invalid_operation_type/fake_account.json",
			ExpInputGetAccountRepositoryFile:        "./testdata/transaction/create/04_should_create_transaction_return_invalid_operation_type/exp_in_get_account_repo.json",
			ExpInputCreateTransactionRepositoryFile: "./testdata/transaction/create/04_should_create_transaction_return_invalid_operation_type/exp_in_create_transaction_repo.json",
			ExpTransactionID:                        "",
			ExpErr:                                  domain.ErrInvalidOperationType,
		},
		{
			Name:                                    "05_should_create_transaction_payment_return_nil_error",
			FakeTransactionFile:                     "./testdata/transaction/create/05_should_create_transaction_payment_return_nil_error/fake_transaction.json",
			FakeTransactionID:                       "123456",
			FakeAccountFile:                         "./testdata/transaction/create/05_should_create_transaction_payment_return_nil_error/fake_account.json",
			ExpInputGetAccountRepositoryFile:        "./testdata/transaction/create/05_should_create_transaction_payment_return_nil_error/exp_in_get_account_repo.json",
			ExpInputCreateTransactionRepositoryFile: "./testdata/transaction/create/05_should_create_transaction_payment_return_nil_error/exp_in_create_transaction_repo.json",
			ExpTransactionID:                        "123456",
			ExpErr:                                  nil,
		},
		{
			Name:                                    "06_should_create_transaction_with_no_accounts_return_not_found_account",
			FakeTransactionFile:                     "./testdata/transaction/create/06_should_create_transaction_with_no_accounts_return_not_found_account/fake_transaction.json",
			FakeTransactionID:                       "123456",
			FakeAccountFile:                         "./testdata/transaction/create/06_should_create_transaction_with_no_accounts_return_not_found_account/fake_account.json",
			ExpInputGetAccountRepositoryFile:        "./testdata/transaction/create/06_should_create_transaction_with_no_accounts_return_not_found_account/exp_in_get_account_repo.json",
			ExpInputCreateTransactionRepositoryFile: "./testdata/transaction/create/06_should_create_transaction_with_no_accounts_return_not_found_account/exp_in_create_transaction_repo.json",
			ExpTransactionID:                        "",
			ExpErr:                                  domain.ErrAccountNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testCreateTransaction(t, tc.Name, tc.FakeTransactionFile, tc.FakeTransactionID, tc.FakeAccountFile, tc.FakeErrCreateTransactionRepository, tc.FakeErrGetAccountRepository, tc.ExpInputGetAccountRepositoryFile, tc.ExpInputCreateTransactionRepositoryFile, tc.ExpTransactionID, tc.ExpErr)
		})
	}
}

func testCreateTransaction(t *testing.T, name, fakeTrFile, fakeTrID, fakeAccFile string, fakeErrCreateTrRepo, fakeErrGetAccRepo error, expInGetAccRepoFile, expInCreateTrRepoFile, expID string, expErr error) {
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
		fakeAcc := testhelpers.ReadJSON[[]domain.Account](t, fakeAccFile)
		return fakeAcc, fakeErrGetAccRepo
	}

	fakeTr := testhelpers.ReadJSON[domain.Transaction](t, fakeTrFile)
	useCase := usecases.NewTransactionUseCase(mockTransaction{}, mockAccount{})
	gotID, gotErr := useCase.Create(context.Background(), fakeTr)
	if *update {
		testhelpers.CreateJSON(t, expInCreateTrRepoFile, gotInCreateTrRepo)
		testhelpers.CreateJSON(t, expInGetAccRepoFile, gotInGetAccRepo)
		return
	}

	assert.Equal(t, expErr, gotErr, "exp error should be equal got error")
	assert.Equal(t, expID, gotID, "exp ID should be equal got ID")
	testhelpers.CompareWithFile(t, "compare input get account repository", expInGetAccRepoFile, gotInGetAccRepo)
	testhelpers.CompareWithFile(t, "compare input create transaction repository", expInCreateTrRepoFile, gotInCreateTrRepo)

}

func TestGetTransactions(t *testing.T) {
	testCases := []struct {
		Name                                  string
		FakeFilterTransactionFile             string
		FakeGetTransactionsRepositoryFile     string
		FakeErrGetTransactionsRepository      error
		LimitByPage                           int64
		ExpInputGetTransactionsRepositoryFile string
		ExpTransactionsPagingFile             string
		ExpErr                                error
	}{
		{
			Name:                                  "01_should_get_transactions_return_nil_error_and_next_page",
			FakeFilterTransactionFile:             "./testdata/transaction/get/01_should_get_transactions_return_nil_error_and_next_page/fake_filter.json",
			FakeGetTransactionsRepositoryFile:     "./testdata/transaction/get/01_should_get_transactions_return_nil_error_and_next_page/fake_get_transactions_repo.json",
			LimitByPage:                           4,
			ExpInputGetTransactionsRepositoryFile: "./testdata/transaction/get/01_should_get_transactions_return_nil_error_and_next_page/exp_in_get_transactions_repo.json",
			ExpTransactionsPagingFile:             "./testdata/transaction/get/01_should_get_transactions_return_nil_error_and_next_page/exp_transactions_paging.json",
			ExpErr:                                nil,
		},
		{
			Name:                                  "02_should_get_transactions_return_nil_error_and_no_next_page",
			FakeFilterTransactionFile:             "./testdata/transaction/get/02_should_get_transactions_return_nil_error_and_no_next_page/fake_filter.json",
			FakeGetTransactionsRepositoryFile:     "./testdata/transaction/get/02_should_get_transactions_return_nil_error_and_no_next_page/fake_get_transactions_repo.json",
			LimitByPage:                           20,
			ExpInputGetTransactionsRepositoryFile: "./testdata/transaction/get/02_should_get_transactions_return_nil_error_and_no_next_page/exp_in_get_transactions_repo.json",
			ExpTransactionsPagingFile:             "./testdata/transaction/get/02_should_get_transactions_return_nil_error_and_no_next_page/exp_transactions_paging.json",
			ExpErr:                                nil,
		},
		{
			Name:                                  "03_should_get_transactions_return_invalid_operation_type",
			FakeFilterTransactionFile:             "./testdata/transaction/get/03_should_get_transactions_return_invalid_operation_type/fake_filter.json",
			FakeGetTransactionsRepositoryFile:     "./testdata/transaction/get/03_should_get_transactions_return_invalid_operation_type/fake_get_transactions_repo.json",
			LimitByPage:                           20,
			ExpInputGetTransactionsRepositoryFile: "./testdata/transaction/get/03_should_get_transactions_return_invalid_operation_type/exp_in_get_transactions_repo.json",
			ExpTransactionsPagingFile:             "./testdata/transaction/get/03_should_get_transactions_return_invalid_operation_type/exp_transactions_paging.json",
			ExpErr:                                domain.ErrInvalidOperationType,
		},
		{
			Name:                                  "04_should_get_transactions_return_unknow_error_on_get_transaction_repo",
			FakeFilterTransactionFile:             "./testdata/transaction/get/04_should_get_transactions_return_unknow_error_on_get_transaction_repo/fake_filter.json",
			FakeGetTransactionsRepositoryFile:     "./testdata/transaction/get/04_should_get_transactions_return_unknow_error_on_get_transaction_repo/fake_get_transactions_repo.json",
			LimitByPage:                           20,
			FakeErrGetTransactionsRepository:      domain.ErrUnknow,
			ExpInputGetTransactionsRepositoryFile: "./testdata/transaction/get/04_should_get_transactions_return_unknow_error_on_get_transaction_repo/exp_in_get_transactions_repo.json",
			ExpTransactionsPagingFile:             "./testdata/transaction/get/04_should_get_transactions_return_unknow_error_on_get_transaction_repo/exp_transactions_paging.json",
			ExpErr:                                domain.ErrUnknow,
		},
		{
			Name:                                  "05_should_get_transactions_return_not_found_from_transactions_paging",
			FakeFilterTransactionFile:             "./testdata/transaction/get/05_should_get_transactions_return_not_found_from_transactions_paging/fake_filter.json",
			FakeGetTransactionsRepositoryFile:     "./testdata/transaction/get/05_should_get_transactions_return_not_found_from_transactions_paging/fake_get_transactions_repo.json",
			LimitByPage:                           20,
			ExpInputGetTransactionsRepositoryFile: "./testdata/transaction/get/05_should_get_transactions_return_not_found_from_transactions_paging/exp_in_get_transactions_repo.json",
			ExpTransactionsPagingFile:             "./testdata/transaction/get/05_should_get_transactions_return_not_found_from_transactions_paging/exp_transactions_paging.json",
			ExpErr:                                domain.ErrTransactionsNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testGetTransactions(t, tc.Name, tc.FakeFilterTransactionFile, tc.FakeGetTransactionsRepositoryFile, tc.LimitByPage, tc.FakeErrGetTransactionsRepository, tc.ExpInputGetTransactionsRepositoryFile, tc.ExpTransactionsPagingFile, tc.ExpErr)
		})
	}
}

func testGetTransactions(t *testing.T, name, fakeFilterTrFile, fakeGetTrsRepoFile string, limitByPage int64, fakeErrGetTrRepo error, expInGetTrRepoFile, expTrsPagFile string, expErr error) {
	domain.LimitByPage = limitByPage
	fakeFilterTr := testhelpers.ReadJSON[domain.TransactionFilter](t, fakeFilterTrFile)
	var gotInGetTrRepo domain.TransactionFilter
	GetTransactionsMock = func(ctx context.Context, tr domain.TransactionFilter) ([]*domain.Transaction, error) {
		gotInGetTrRepo = tr
		return testhelpers.ReadJSON[[]*domain.Transaction](t, fakeGetTrsRepoFile), fakeErrGetTrRepo
	}

	useCase := usecases.NewTransactionUseCase(mockTransaction{}, mockAccount{})
	gotTrsPag, gotErr := useCase.Get(context.Background(), fakeFilterTr)
	if *update {
		testhelpers.CreateJSON(t, expTrsPagFile, gotTrsPag)
		testhelpers.CreateJSON(t, expInGetTrRepoFile, gotInGetTrRepo)
		return
	}

	assert.Equal(t, expErr, gotErr, "exp error should be equal got error")
	testhelpers.CompareWithFile(t, "compare input get transaction repository", expInGetTrRepoFile, gotInGetTrRepo)
	testhelpers.CompareWithFile(t, "compare return get transactions paging", expTrsPagFile, gotTrsPag)
}
