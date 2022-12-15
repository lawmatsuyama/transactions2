package apimanager_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/lawmatsuyama/pismo-transactions/domain"
	"github.com/lawmatsuyama/pismo-transactions/infra/apimanager"
	"github.com/lawmatsuyama/pismo-transactions/pkg/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	testCases := []struct {
		Name                             string
		FakeCreateAccountRequestFile     string
		FakeIDCreateAccountUseCase       string
		FakeErrCreateAccountUseCase      error
		FakeErrResponseWrite             error
		ExpInputCreateAccountUseCaseFile string
		ExpResponseFile                  string
		ExpStatusCode                    int
	}{
		{
			Name:                             "01_should_create_account_return_nil_error",
			FakeCreateAccountRequestFile:     "./testdata/account/create/01_should_create_account_return_nil_error/request.json",
			FakeIDCreateAccountUseCase:       "123456",
			ExpInputCreateAccountUseCaseFile: "./testdata/account/create/01_should_create_account_return_nil_error/exp_in_create_acc_usecase.json",
			ExpResponseFile:                  "./testdata/account/create/01_should_create_account_return_nil_error/exp_response.json",
			ExpStatusCode:                    http.StatusOK,
		},
		{
			Name:                             "02_should_create_account_return_error_on_decode",
			FakeCreateAccountRequestFile:     "./testdata/account/create/02_should_create_account_return_error_on_decode/request.json",
			FakeIDCreateAccountUseCase:       "123456",
			ExpInputCreateAccountUseCaseFile: "./testdata/account/create/02_should_create_account_return_error_on_decode/exp_in_create_acc_usecase.json",
			ExpResponseFile:                  "./testdata/account/create/02_should_create_account_return_error_on_decode/exp_response.json",
			ExpStatusCode:                    http.StatusBadRequest,
		},
		{
			Name:                             "03_should_create_account_return_error_create_account_usecase",
			FakeCreateAccountRequestFile:     "./testdata/account/create/03_should_create_account_return_error_create_account_usecase/request.json",
			FakeIDCreateAccountUseCase:       "123456",
			FakeErrCreateAccountUseCase:      domain.ErrUnknown,
			ExpInputCreateAccountUseCaseFile: "./testdata/account/create/03_should_create_account_return_error_create_account_usecase/exp_in_create_acc_usecase.json",
			ExpResponseFile:                  "./testdata/account/create/03_should_create_account_return_error_create_account_usecase/exp_response.json",
			ExpStatusCode:                    http.StatusBadRequest,
		},
		{
			Name:                             "04_should_create_account_return_error_on_write_response",
			FakeCreateAccountRequestFile:     "./testdata/account/create/04_should_create_account_return_error_on_write_response/request.json",
			FakeIDCreateAccountUseCase:       "123456",
			FakeErrResponseWrite:             domain.ErrUnknown,
			ExpInputCreateAccountUseCaseFile: "./testdata/account/create/04_should_create_account_return_error_on_write_response/exp_in_create_acc_usecase.json",
			ExpResponseFile:                  "./testdata/account/create/04_should_create_account_return_error_on_write_response/exp_response.json",
			ExpStatusCode:                    http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testCreateAccount(t, tc.Name, tc.FakeCreateAccountRequestFile, tc.FakeIDCreateAccountUseCase, tc.FakeErrCreateAccountUseCase, tc.FakeErrResponseWrite, tc.ExpInputCreateAccountUseCaseFile, tc.ExpResponseFile, tc.ExpStatusCode)
		})
	}
}

func testCreateAccount(t *testing.T,
	name, fakeCreateAccReqFile, fakeIDCreateAccUseCase string,
	fakeErrCreateAccUseCase, fakeErrRespWrite error,
	expInCreateAccUseCaseFile, expRespFile string,
	expStatusCode int) {

	var gotInCreateAccUseCase domain.Account
	CreateAccountMock = func(ctx context.Context, acc domain.Account) (id string, err error) {
		gotInCreateAccUseCase = acc
		return fakeIDCreateAccUseCase, fakeErrCreateAccUseCase
	}

	var gotStatusCode int
	WriteHeaderMock = func(statusCode int) {
		gotStatusCode = statusCode
	}

	var gotResp apimanager.GenericResponse[any]
	WriteMock = func(b []byte) (int, error) {
		err := json.Unmarshal(b, &gotResp)
		if err != nil {
			return 0, nil
		}

		return len(b), fakeErrRespWrite
	}

	req := testhelpers.RequestPost(t, "POST", fakeCreateAccReqFile)
	accAPI := apimanager.NewAccountAPI(mockAccount{})
	accAPI.Create(mockResponseWriter{}, req)

	if *update {
		testhelpers.CreateJSON(t, expInCreateAccUseCaseFile, gotInCreateAccUseCase)
		testhelpers.CreateJSON(t, expRespFile, gotResp)
		return
	}

	assert.Equal(t, expStatusCode, gotStatusCode, "exp status code should be equal got status code")
	testhelpers.CompareWithFile(t, "compare input create account use case", expInCreateAccUseCaseFile, gotInCreateAccUseCase)
	testhelpers.CompareWithFile(t, "compare response create account", expRespFile, gotResp)
}

func TestGetAccount(t *testing.T) {
	testCases := []struct {
		Name                           string
		FakeAccountID                  string
		FakeGetAccountsUseCaseFile     string
		FakeErrGetAccountsUseCase      error
		FakeErrResponseWrite           error
		ExpInputGetAccountsUseCaseFile string
		ExpResponseFile                string
		ExpStatusCode                  int
	}{
		{
			Name:                           "01_should_get_account_return_nil_error",
			FakeAccountID:                  "123456",
			FakeGetAccountsUseCaseFile:     "./testdata/account/get/01_should_get_account_return_nil_error/fake_get_account_use_case.json",
			ExpInputGetAccountsUseCaseFile: "./testdata/account/get/01_should_get_account_return_nil_error/exp_in_get_account_use_case.json",
			ExpResponseFile:                "./testdata/account/get/01_should_get_account_return_nil_error/exp_response.json",
			ExpStatusCode:                  http.StatusOK,
		},
		{
			Name:                           "02_should_get_account_return_error_on_get_account_use_case",
			FakeAccountID:                  "123456",
			FakeGetAccountsUseCaseFile:     "./testdata/account/get/02_should_get_account_return_error_on_get_account_use_case/fake_get_account_use_case.json",
			FakeErrGetAccountsUseCase:      domain.ErrUnknown,
			ExpInputGetAccountsUseCaseFile: "./testdata/account/get/02_should_get_account_return_error_on_get_account_use_case/exp_in_get_account_use_case.json",
			ExpResponseFile:                "./testdata/account/get/02_should_get_account_return_error_on_get_account_use_case/exp_response.json",
			ExpStatusCode:                  http.StatusBadRequest,
		},
		{
			Name:                           "03_should_get_account_return_error_not_found",
			FakeAccountID:                  "123456",
			FakeGetAccountsUseCaseFile:     "./testdata/account/get/03_should_get_account_return_error_not_found/fake_get_account_use_case.json",
			ExpInputGetAccountsUseCaseFile: "./testdata/account/get/03_should_get_account_return_error_not_found/exp_in_get_account_use_case.json",
			ExpResponseFile:                "./testdata/account/get/03_should_get_account_return_error_not_found/exp_response.json",
			ExpStatusCode:                  http.StatusNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testGetAccount(t, tc.Name, tc.FakeAccountID, tc.FakeGetAccountsUseCaseFile, tc.FakeErrGetAccountsUseCase, tc.FakeErrResponseWrite, tc.ExpInputGetAccountsUseCaseFile, tc.ExpResponseFile, tc.ExpStatusCode)
		})
	}
}

func testGetAccount(t *testing.T, name, fakeAccID, fakeGetAccsUseCaseFile string,
	fakeErrGetAccsUseCase, fakeErrRespWrite error,
	expInGetAccsUseCaseFile, expRespFile string,
	expStatusCode int) {

	var gotInGetAccsUseCase domain.AccountFilter
	GetAccountMock = func(ctx context.Context, filter domain.AccountFilter) ([]domain.Account, error) {
		gotInGetAccsUseCase = filter
		fakeGetAccsUseCase := testhelpers.ReadJSON[[]domain.Account](t, fakeGetAccsUseCaseFile)
		return fakeGetAccsUseCase, fakeErrGetAccsUseCase
	}

	var gotStatusCode int
	WriteHeaderMock = func(statusCode int) {
		gotStatusCode = statusCode
	}

	var gotResp apimanager.GenericResponse[any]
	WriteMock = func(b []byte) (int, error) {
		err := json.Unmarshal(b, &gotResp)
		if err != nil {
			return 0, nil
		}

		return len(b), fakeErrRespWrite
	}

	req := httptest.NewRequest("GET", "http://localhost/account/{accountID}", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("accountID", fakeAccID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	accAPI := apimanager.NewAccountAPI(mockAccount{})
	accAPI.GetByID(mockResponseWriter{}, req)

	if *update {
		testhelpers.CreateJSON(t, expInGetAccsUseCaseFile, gotInGetAccsUseCase)
		testhelpers.CreateJSON(t, expRespFile, gotResp)
		return
	}

	assert.Equal(t, expStatusCode, gotStatusCode, "exp status code should be equal got status code")
	testhelpers.CompareWithFile(t, "compare input get account use case", expInGetAccsUseCaseFile, gotInGetAccsUseCase)
	testhelpers.CompareWithFile(t, "compare response get account", expRespFile, gotResp)

}
