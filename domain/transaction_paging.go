package domain

// TransactionsPaging represents transactions by page
type TransactionsPaging struct {
	Transactions []*Transaction `json:"transactions"`
	Paging       *Paging        `json:"paging,omitempty"`
}

// NewTransactionsPage returns a new TransactionsPaging
func NewTransactionsPage(trs []*Transaction, pg *Paging) TransactionsPaging {
	return TransactionsPaging{Transactions: trs, Paging: pg}
}

// IsValid check if TransactionsPaging is valid
func (trs TransactionsPaging) IsValid() error {
	if len(trs.Transactions) == 0 {
		return ErrTransactionsNotFound
	}

	return nil
}

// SetNextPaging set the next paging according to number of transactions
func (trs *TransactionsPaging) SetNextPaging() {
	if trs == nil {
		return
	}

	if trs.Paging == nil {
		trs.Paging = &Paging{}
	}

	trs.Paging.SetNextPage(len(trs.Transactions))
}
