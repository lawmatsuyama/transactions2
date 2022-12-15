package apimanager_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/lawmatsuyama/pismo-transactions/domain"
	"github.com/lawmatsuyama/pismo-transactions/infra/apimanager"
	"github.com/lawmatsuyama/pismo-transactions/pkg/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	testCases := []struct {
		Name                                 string
		FakeCreateTransactionRequestFile     string
		FakeIDCreateTransactionUseCase       string
		FakeErrCreateTransactionUseCase      error
		FakeErrResponseWrite                 error
		ExpInputCreateTransactionUseCaseFile string
		ExpResponseFile                      string
		ExpStatusCode                        int
	}{
		{
			Name:                                 "01_should_create_transaction_return_nil_error",
			FakeCreateTransactionRequestFile:     "./testdata/transaction/create/01_should_create_transaction_return_nil_error/request.json",
			FakeIDCreateTransactionUseCase:       "123456",
			ExpInputCreateTransactionUseCaseFile: "./testdata/transaction/create/01_should_create_transaction_return_nil_error/exp_in_create_tr_usecase.json",
			ExpResponseFile:                      "./testdata/transaction/create/01_should_create_transaction_return_nil_error/exp_response.json",
			ExpStatusCode:                        http.StatusOK,
		},
		{
			Name:                                 "02_should_create_transaction_return_error_on_decode",
			FakeCreateTransactionRequestFile:     "./testdata/transaction/create/02_should_create_transaction_return_error_on_decode/request.json",
			FakeIDCreateTransactionUseCase:       "123456",
			ExpInputCreateTransactionUseCaseFile: "./testdata/transaction/create/02_should_create_transaction_return_error_on_decode/exp_in_create_tr_usecase.json",
			ExpResponseFile:                      "./testdata/transaction/create/02_should_create_transaction_return_error_on_decode/exp_response.json",
			ExpStatusCode:                        http.StatusBadRequest,
		},
		{
			Name:                                 "03_should_create_transaction_return_error_create_transaction_usecase",
			FakeCreateTransactionRequestFile:     "./testdata/transaction/create/03_should_create_transaction_return_error_create_transaction_usecase/request.json",
			FakeIDCreateTransactionUseCase:       "123456",
			FakeErrCreateTransactionUseCase:      domain.ErrUnknow,
			ExpInputCreateTransactionUseCaseFile: "./testdata/transaction/create/03_should_create_transaction_return_error_create_transaction_usecase/exp_in_create_tr_usecase.json",
			ExpResponseFile:                      "./testdata/transaction/create/03_should_create_transaction_return_error_create_transaction_usecase/exp_response.json",
			ExpStatusCode:                        http.StatusBadRequest,
		},
		{
			Name:                                 "04_should_create_transaction_return_error_on_write_response",
			FakeCreateTransactionRequestFile:     "./testdata/transaction/create/04_should_create_transaction_return_error_on_write_response/request.json",
			FakeIDCreateTransactionUseCase:       "123456",
			FakeErrResponseWrite:                 domain.ErrUnknow,
			ExpInputCreateTransactionUseCaseFile: "./testdata/transaction/create/04_should_create_transaction_return_error_on_write_response/exp_in_create_tr_usecase.json",
			ExpResponseFile:                      "./testdata/transaction/create/04_should_create_transaction_return_error_on_write_response/exp_response.json",
			ExpStatusCode:                        http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testCreateTransaction(t, tc.Name, tc.FakeCreateTransactionRequestFile, tc.FakeIDCreateTransactionUseCase, tc.FakeErrCreateTransactionUseCase, tc.FakeErrResponseWrite, tc.ExpInputCreateTransactionUseCaseFile, tc.ExpResponseFile, tc.ExpStatusCode)
		})
	}
}

func testCreateTransaction(t *testing.T,
	name, fakeCreateTrReqFile, fakeIDCreateTrUseCase string,
	fakeErrCreateTrUseCase, fakeErrRespWrite error,
	expInCreateTrUseCaseFile, expRespFile string,
	expStatusCode int) {

	var gotInCreateTrUseCase domain.Transaction
	CreateTransactionMock = func(ctx context.Context, tr domain.Transaction) (id string, err error) {
		gotInCreateTrUseCase = tr
		return fakeIDCreateTrUseCase, fakeErrCreateTrUseCase
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

	req := testhelpers.RequestPost(t, "POST", fakeCreateTrReqFile)
	trAPI := apimanager.NewTransactionAPI(mockTransaction{})
	trAPI.Create(mockResponseWriter{}, req)

	if *update {
		testhelpers.CreateJSON(t, expInCreateTrUseCaseFile, gotInCreateTrUseCase)
		testhelpers.CreateJSON(t, expRespFile, gotResp)
		return
	}

	assert.Equal(t, expStatusCode, gotStatusCode, "exp status code should be equal got status code")
	testhelpers.CompareWithFile(t, "compare input create transaction use case", expInCreateTrUseCaseFile, gotInCreateTrUseCase)
	testhelpers.CompareWithFile(t, "compare response create transaction", expRespFile, gotResp)
}

func TestGetTransaction(t *testing.T) {
	testCases := []struct {
		Name                              string
		FakeGetTransactionRequestFile     string
		FakeGetTransactionUseCaseFile     string
		FakeErrGetTransactionUseCase      error
		FakeErrResponseWrite              error
		ExpInputGetTransactionUseCaseFile string
		ExpResponseFile                   string
		ExpStatusCode                     int
	}{
		{
			Name:                              "01_should_get_transaction_return_nil_error",
			FakeGetTransactionRequestFile:     "./testdata/transaction/get/01_should_get_transaction_return_nil_error/request.json",
			FakeGetTransactionUseCaseFile:     "./testdata/transaction/get/01_should_get_transaction_return_nil_error/fake_get_transaction_use_case.json",
			ExpInputGetTransactionUseCaseFile: "./testdata/transaction/get/01_should_get_transaction_return_nil_error/exp_in_get_tr_usecase.json",
			ExpResponseFile:                   "./testdata/transaction/get/01_should_get_transaction_return_nil_error/exp_response.json",
			ExpStatusCode:                     http.StatusOK,
		},
		{
			Name:                              "02_should_get_transaction_return_error_on_decode",
			FakeGetTransactionRequestFile:     "./testdata/transaction/get/02_should_get_transaction_return_error_on_decode/request.json",
			FakeGetTransactionUseCaseFile:     "./testdata/transaction/get/02_should_get_transaction_return_error_on_decode/fake_get_transaction_use_case.json",
			ExpInputGetTransactionUseCaseFile: "./testdata/transaction/get/02_should_get_transaction_return_error_on_decode/exp_in_get_tr_usecase.json",
			ExpResponseFile:                   "./testdata/transaction/get/02_should_get_transaction_return_error_on_decode/exp_response.json",
			ExpStatusCode:                     http.StatusBadRequest,
		},
		{
			Name:                              "03_should_get_transaction_return_error_get_transaction_usecase",
			FakeGetTransactionRequestFile:     "./testdata/transaction/get/03_should_get_transaction_return_error_get_transaction_usecase/request.json",
			FakeGetTransactionUseCaseFile:     "./testdata/transaction/get/03_should_get_transaction_return_error_get_transaction_usecase/fake_get_transaction_use_case.json",
			FakeErrGetTransactionUseCase:      domain.ErrUnknow,
			ExpInputGetTransactionUseCaseFile: "./testdata/transaction/get/03_should_get_transaction_return_error_get_transaction_usecase/exp_in_get_tr_usecase.json",
			ExpResponseFile:                   "./testdata/transaction/get/03_should_get_transaction_return_error_get_transaction_usecase/exp_response.json",
			ExpStatusCode:                     http.StatusBadRequest,
		},
		{
			Name:                              "04_should_get_transaction_return_error_on_write_response",
			FakeGetTransactionRequestFile:     "./testdata/transaction/get/04_should_get_transaction_return_error_on_write_response/request.json",
			FakeGetTransactionUseCaseFile:     "./testdata/transaction/get/04_should_get_transaction_return_error_on_write_response/fake_get_transaction_use_case.json",
			FakeErrResponseWrite:              domain.ErrUnknow,
			ExpInputGetTransactionUseCaseFile: "./testdata/transaction/get/04_should_get_transaction_return_error_on_write_response/exp_in_get_tr_usecase.json",
			ExpResponseFile:                   "./testdata/transaction/get/04_should_get_transaction_return_error_on_write_response/exp_response.json",
			ExpStatusCode:                     http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testGetTransaction(t, tc.Name, tc.FakeGetTransactionRequestFile, tc.FakeGetTransactionUseCaseFile, tc.FakeErrGetTransactionUseCase, tc.FakeErrResponseWrite, tc.ExpInputGetTransactionUseCaseFile, tc.ExpResponseFile, tc.ExpStatusCode)
		})
	}
}

func testGetTransaction(t *testing.T,
	name, fakeGetTrReqFile, fakeGetTrUseCaseFile string,
	fakeErrGetTrUseCase, fakeErrRespWrite error,
	expInGetTrUseCaseFile, expRespFile string,
	expStatusCode int) {

	var gotInGetTrUseCase domain.TransactionFilter
	GetTransactionMock = func(ctx context.Context, tr domain.TransactionFilter) (trPag domain.TransactionsPaging, err error) {
		gotInGetTrUseCase = tr
		fakeGetTrUseCase := testhelpers.ReadJSON[domain.TransactionsPaging](t, fakeGetTrUseCaseFile)
		return fakeGetTrUseCase, fakeErrGetTrUseCase
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

	req := testhelpers.RequestPost(t, "POST", fakeGetTrReqFile)
	trAPI := apimanager.NewTransactionAPI(mockTransaction{})
	trAPI.Get(mockResponseWriter{}, req)

	if *update {
		testhelpers.CreateJSON(t, expInGetTrUseCaseFile, gotInGetTrUseCase)
		testhelpers.CreateJSON(t, expRespFile, gotResp)
		return
	}

	assert.Equal(t, expStatusCode, gotStatusCode, "exp status code should be equal got status code")
	testhelpers.CompareWithFile(t, "compare input create transaction use case", expInGetTrUseCaseFile, gotInGetTrUseCase)
	testhelpers.CompareWithFile(t, "compare response create transaction", expRespFile, gotResp)
}
