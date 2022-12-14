package domain

import "context"

type TransactionValidator interface {
	IsValid() error
}

// UUIDGenerator represents an UUID generator
type UUIDGenerator interface {
	Generate() string
}

type AccountRepository interface {
	Create(ctx context.Context, acc Account) (err error)
	Get(ctx context.Context, filter AccountFilter) (accs []Account, err error)
}

type AccountUseCase interface {
	Create(ctx context.Context, acc Account) (id string, err error)
	Get(ctx context.Context, filter AccountFilter) (accs []Account, err error)
}

type TransactionRepository interface {
	Create(ctx context.Context, tr Transaction) (err error)
	// Get(ctx context.Context, tr TransactionFilter) ([]Transaction, error)
}

type TransactionUseCase interface {
	Create(ctx context.Context, tr Transaction) (id string, err error)
}
