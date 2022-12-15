package domain

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

// DocumentNumber represents an user document number
type DocumentNumber string

// String returns document number in string type
func (doc DocumentNumber) String() string {
	return string(doc)
}

// IsValid returns error if document number is invalid
func (doc DocumentNumber) IsValid() error {
	if len(doc) != 11 {
		return ErrInvalidDocumentNumber
	}

	n, err := strconv.Atoi(doc.String())
	if err != nil {
		return ErrInvalidDocumentNumber
	}

	digits := n % 100
	n = n / 100

	sum1, sum2 := 0, 0
	allEqual := true
	first := true
	previous := 0
	for i := 9; i > 0; i-- {
		remainder := n % 10
		n = n / 10
		if !first {
			allEqual = allEqual && remainder == previous
		} else {
			first = false
		}
		sum1 = sum1 + (remainder * i)
		sum2 = sum2 + (remainder * (i - 1))
		previous = remainder
	}

	if allEqual {
		logrus.Debug("invalid document number - all equal digits")
		return ErrInvalidDocumentNumber
	}

	digitCalculated1 := remainderDigit(sum1, 11)
	digitCalculated2 := remainderDigit(digitCalculated1*9+sum2, 11)

	digitReceived1 := (digits / 10) % 10
	digitReceived2 := digits % 10

	if digitCalculated1 != digitReceived1 ||
		digitCalculated2 != digitReceived2 {
		logrus.Debug("invalid document number - digits are invalid")
		return ErrInvalidDocumentNumber
	}

	return nil
}

func remainderDigit(sum, divisor int) int {
	digit := sum % divisor
	if digit == 10 {
		return 0
	}
	return digit
}
