package domain

var (
	// LimitByPage is a global variable to limit the transactions of listing page
	LimitByPage int64 = 20
)

// Paging represents paging struct
type Paging struct {
	Page     int64  `json:"page"`
	NextPage *int64 `json:"next_page,omitempty"`
}

// CurrentPage returns the current page number
func (p *Paging) CurrentPage() int64 {
	if p == nil || p.Page <= 0 {
		return 1
	}

	return p.Page
}

// SetNextPage set next page according to num transactions
func (p *Paging) SetNextPage(numTransactions int) {
	if p == nil {
		return
	}

	if p.Page <= 0 {
		p.Page = 1
	}

	if int64(numTransactions) >= LimitByPage {
		nextPage := p.Page + 1
		p.NextPage = &nextPage
	}
}

// Skip returns the number of documents to skip to view the page
func (p *Paging) Skip() int64 {
	page := p.CurrentPage()
	return (page - 1) * LimitByPage
}

// LimitByPage returns the limit number of documents by page
func (p *Paging) LimitByPage() int64 {
	return LimitByPage
}
