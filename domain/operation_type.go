package domain

var mapOperationType = map[OperationType]OperationTypeDetails{
	1: NewOperationTypeDetails("COMPRA A VISTA", -1.0),
	2: NewOperationTypeDetails("COMPRA PARCELADA", -1.0),
	3: NewOperationTypeDetails("SAQUE", -1.0),
	4: NewOperationTypeDetails("PAGAMENTO", 1.0),
}

type OperationType int64

type OperationTypeDetails struct {
	Description string
	AmountSign  float64
}

func NewOperationTypeDetails(desc string, sign float64) OperationTypeDetails {
	return OperationTypeDetails{Description: desc, AmountSign: sign}
}

func (oper OperationType) String() string {
	operDetails, ok := mapOperationType[oper]
	if !ok {
		return ""
	}

	return operDetails.Description
}

func (oper OperationType) IsValid() error {
	if _, ok := mapOperationType[oper]; !ok {
		return ErrInvalidOperationType
	}

	return nil
}

func (oper OperationType) Sign() float64 {
	operDetails, ok := mapOperationType[oper]
	if !ok {
		return 1.0
	}

	return operDetails.AmountSign
}
