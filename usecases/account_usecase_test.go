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

func TestCreateAccount(t *testing.T) {
	testCases := []struct {
		Name                         string
		FakeAccountFile              string
		FakeErrCreateRepository      error
		FakeAccountID                string
		ExpInputCreateRepositoryFile string
		ExpID                        string
		ExpError                     error
	}{
		{
			Name:                         "01_should_create_account_return_error_nil",
			FakeAccountFile:              "./testdata/account/create/01_should_create_account_return_error_nil/fake_account.json",
			FakeAccountID:                "123456",
			ExpInputCreateRepositoryFile: "./testdata/account/create/01_should_create_account_return_error_nil/exp_in_create_repository.json",
			ExpID:                        "123456",
			ExpError:                     nil,
		},
		{
			Name:                         "02_should_create_account_return_invalid_document_error",
			FakeAccountFile:              "./testdata/account/create/02_should_create_account_return_invalid_document_error/fake_account.json",
			ExpInputCreateRepositoryFile: "./testdata/account/create/02_should_create_account_return_invalid_document_error/exp_in_create_repository.json",
			ExpID:                        "",
			ExpError:                     domain.ErrInvalidDocumentNumber,
		},
		{
			Name:                         "03_should_create_account_return_error_on_repository",
			FakeAccountFile:              "./testdata/account/create/03_should_create_account_return_error_on_repository/fake_account.json",
			FakeAccountID:                "999999",
			ExpInputCreateRepositoryFile: "./testdata/account/create/03_should_create_account_return_error_on_repository/exp_in_create_repository.json",
			FakeErrCreateRepository:      domain.ErrUnknown,
			ExpID:                        "",
			ExpError:                     domain.ErrUnknown,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testCreateAccount(t, tc.Name, tc.FakeAccountFile, tc.FakeAccountID, tc.FakeErrCreateRepository, tc.ExpInputCreateRepositoryFile, tc.ExpID, tc.ExpError)
		})
	}
}

func testCreateAccount(t *testing.T, name, fakeAccFile, fakeAccID string, fakeErrCreateRepo error, expInCreateRepoFile, expID string, expErr error) {
	domain.Now = func() time.Time {
		return domain.TimeSaoPaulo(time.Date(2022, 10, 10, 12, 0, 0, 0, time.UTC))
	}

	domain.UUID = mockUUID{}
	UUIDGenerateMock = func() string {
		return fakeAccID
	}

	fakeAcc := testhelpers.ReadJSON[domain.Account](t, fakeAccFile)
	var gotInCreateRepo domain.Account
	CreateAccountMock = func(ctx context.Context, acc domain.Account) (err error) {
		gotInCreateRepo = acc
		return fakeErrCreateRepo
	}

	useCase := usecases.NewAccountUseCase(mockAccount{})
	gotID, gotErr := useCase.Create(context.Background(), fakeAcc)
	if *update {
		testhelpers.CreateJSON(t, expInCreateRepoFile, gotInCreateRepo)
		return
	}

	assert.Equal(t, expErr, gotErr, "exp error should be equal got error")
	assert.Equal(t, expID, gotID, "exp ID should be equal got ID")
	testhelpers.CompareWithFile(t, "compare input create account repository", expInCreateRepoFile, gotInCreateRepo)
}

func TestGetAccounts(t *testing.T) {
	testCases := []struct {
		Name                             string
		FakeFilterFile                   string
		FakeGetAccountsRepositoryFile    string
		FakeErrGetAccountRepository      error
		ExpInputGetAccountRepositoryFile string
		ExpAccountsFile                  string
		ExpError                         error
	}{
		{
			Name:                             "01_should_get_accounts_return_error_nil",
			FakeFilterFile:                   "./testdata/account/get/01_should_get_accounts_return_error_nil/fake_filter.json",
			FakeGetAccountsRepositoryFile:    "./testdata/account/get/01_should_get_accounts_return_error_nil/fake_get_accounts.json",
			ExpInputGetAccountRepositoryFile: "./testdata/account/get/01_should_get_accounts_return_error_nil/exp_in_get_account_repository.json",
			ExpAccountsFile:                  "./testdata/account/get/01_should_get_accounts_return_error_nil/exp_accounts.json",
		},
		{
			Name:                             "02_should_get_accounts_return_error_invalid_document_number",
			FakeFilterFile:                   "./testdata/account/get/02_should_get_accounts_return_error_invalid_document_number/fake_filter.json",
			FakeGetAccountsRepositoryFile:    "./testdata/account/get/02_should_get_accounts_return_error_invalid_document_number/fake_get_accounts.json",
			ExpInputGetAccountRepositoryFile: "./testdata/account/get/02_should_get_accounts_return_error_invalid_document_number/exp_in_get_account_repository.json",
			ExpAccountsFile:                  "./testdata/account/get/02_should_get_accounts_return_error_invalid_document_number/exp_accounts.json",
			ExpError:                         domain.ErrInvalidDocumentNumber,
		},
		{
			Name:                             "03_should_get_accounts_return_error_on_get_account_repository",
			FakeFilterFile:                   "./testdata/account/get/03_should_get_accounts_return_error_on_get_account_repository/fake_filter.json",
			FakeGetAccountsRepositoryFile:    "./testdata/account/get/03_should_get_accounts_return_error_on_get_account_repository/fake_get_accounts.json",
			FakeErrGetAccountRepository:      domain.ErrUnknown,
			ExpInputGetAccountRepositoryFile: "./testdata/account/get/03_should_get_accounts_return_error_on_get_account_repository/exp_in_get_account_repository.json",
			ExpAccountsFile:                  "./testdata/account/get/03_should_get_accounts_return_error_on_get_account_repository/exp_accounts.json",
			ExpError:                         domain.ErrUnknown,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testGetAccounts(t, tc.Name, tc.FakeFilterFile, tc.FakeGetAccountsRepositoryFile, tc.ExpInputGetAccountRepositoryFile, tc.ExpAccountsFile, tc.FakeErrGetAccountRepository, tc.ExpError)
		})
	}
}

func testGetAccounts(t *testing.T, name, fakeFilterFile, fakeGetAccsRepoFile, expInGetAccRepoFile, expAccsFile string, fakeErrGetAccRepo, expErr error) {
	fakeFilter := testhelpers.ReadJSON[domain.AccountFilter](t, fakeFilterFile)
	var gotInGetAccRepo domain.AccountFilter
	GetAccountMock = func(ctx context.Context, filter domain.AccountFilter) (accs []domain.Account, err error) {
		gotInGetAccRepo = filter
		fakeGetAccRepo := testhelpers.ReadJSON[[]domain.Account](t, fakeGetAccsRepoFile)
		return fakeGetAccRepo, fakeErrGetAccRepo
	}

	useCase := usecases.NewAccountUseCase(mockAccount{})
	gotAccs, gotErr := useCase.Get(context.Background(), fakeFilter)

	if *update {
		testhelpers.CreateJSON(t, expInGetAccRepoFile, gotInGetAccRepo)
		testhelpers.CreateJSON(t, expAccsFile, gotAccs)
		return
	}

	assert.Equal(t, expErr, gotErr, "exp error should be equal got error")
	testhelpers.CompareWithFile(t, "compare input get create account repository", expInGetAccRepoFile, gotInGetAccRepo)
	testhelpers.CompareWithFile(t, "compare return get create account repository", expAccsFile, gotAccs)
}
