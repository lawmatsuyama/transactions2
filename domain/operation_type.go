package domain

var mapOperationType = map[OperationType]OperationTypeDetails{
	1: NewOperationTypeDetails("COMPRA A VISTA", -1.0),
	2: NewOperationTypeDetails("COMPRA PARCELADA", -1.0),
	3: NewOperationTypeDetails("SAQUE", -1.0),
	4: NewOperationTypeDetails("PAGAMENTO", 1.0),
}

// OperationType represents an operation type of transaction
type OperationType int64

// OperationTypeDetails contains details about operation type
type OperationTypeDetails struct {
	Description string
	AmountSign  float64
}

// NewOperationTypeDetails returns a new OperationTypeDetails
func NewOperationTypeDetails(desc string, sign float64) OperationTypeDetails {
	return OperationTypeDetails{Description: desc, AmountSign: sign}
}

// String returns OperationType in string type
func (oper OperationType) String() string {
	operDetails, ok := mapOperationType[oper]
	if !ok {
		return ""
	}

	return operDetails.Description
}

// IsValid returns error if operation type is invalid
func (oper OperationType) IsValid() error {
	if _, ok := mapOperationType[oper]; !ok {
		return ErrInvalidOperationType
	}

	return nil
}

// Sign returns the math sign related to the operation type
func (oper OperationType) Sign() float64 {
	operDetails, ok := mapOperationType[oper]
	if !ok {
		return 1.0
	}

	return operDetails.AmountSign
}
