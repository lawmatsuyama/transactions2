package domain_test

import (
	"testing"

	"github.com/lawmatsuyama/pismo-transactions/domain"
	"github.com/stretchr/testify/assert"
)

func TestOperationTypeString(t *testing.T) {
	t.Run("01_should_return_valid_string", func(t *testing.T) {
		oper := domain.OperationType(1)
		got := oper.String()
		exp := "COMPRA A VISTA"

		assert.Equal(t, exp, got)
	})

	t.Run("02_should_return_empty_string", func(t *testing.T) {
		oper := domain.OperationType(99)
		got := oper.String()
		exp := ""

		assert.Equal(t, exp, got)
	})
}

func TestOperationTypeIsValid(t *testing.T) {
	t.Run("01_should_return_nil_error", func(t *testing.T) {
		oper := domain.OperationType(1)
		got := oper.IsValid()
		exp := error(nil)
		assert.Equal(t, exp, got)
	})

	t.Run("02_should_return_invalid_operation_type_error", func(t *testing.T) {
		oper := domain.OperationType(99)
		got := oper.IsValid()
		exp := domain.ErrInvalidOperationType

		assert.Equal(t, exp, got)
	})
}

func TestOperationTypeSign(t *testing.T) {
	t.Run("01_should_return_valid_sign", func(t *testing.T) {
		oper := domain.OperationType(3)
		got := oper.Sign()
		exp := -1.0
		assert.Equal(t, exp, got)
	})

	t.Run("02_should_return_default_sign", func(t *testing.T) {
		oper := domain.OperationType(99)
		got := oper.Sign()
		exp := 1.0

		assert.Equal(t, exp, got)
	})
}
