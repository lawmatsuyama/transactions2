package domain_test

import (
	"testing"

	"github.com/lawmatsuyama/pismo-transactions/domain"
	"github.com/stretchr/testify/assert"
)

func TestDocumentNumberIsValid(t *testing.T) {
	testCases := []struct {
		Name           string
		DocumentNumber string
		ExpectedError  error
	}{
		{
			Name:           "01_should_return_valid_document",
			DocumentNumber: "32877847900",
			ExpectedError:  nil,
		},
		{
			Name:           "02_should_return_invalid_document_all_equals",
			DocumentNumber: "44444444444",
			ExpectedError:  domain.ErrInvalidDocumentNumber,
		},
		{
			Name:           "03_should_return_invalid_document",
			DocumentNumber: "12345678910",
			ExpectedError:  domain.ErrInvalidDocumentNumber,
		},
		{
			Name:           "04_should_return_invalid_document_more_than_11_digits",
			DocumentNumber: "123456789101145",
			ExpectedError:  domain.ErrInvalidDocumentNumber,
		},
		{
			Name:           "05_should_return_invalid_document_with_string",
			DocumentNumber: "12345678S10",
			ExpectedError:  domain.ErrInvalidDocumentNumber,
		},
		{
			Name:           "06_should_return_invalid_document_with_empty",
			DocumentNumber: "",
			ExpectedError:  domain.ErrInvalidDocumentNumber,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testDocumentNumberIsValid(t, tc.Name, tc.DocumentNumber, tc.ExpectedError)
		})
	}
}

func testDocumentNumberIsValid(t *testing.T, name, docNumStr string, expErr error) {
	docNum := domain.DocumentNumber(docNumStr)
	err := docNum.IsValid()
	assert.Equal(t, expErr, err)
}
