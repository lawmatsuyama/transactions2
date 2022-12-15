package domain

import "context"

// TransactionValidator represents a validator to check transaction data
type TransactionValidator interface {
	IsValid() error
}

// UUIDGenerator represents an UUID generator
type UUIDGenerator interface {
	Generate() string
}

// AccountRepository represents an account repository
type AccountRepository interface {
	Create(ctx context.Context, acc Account) (err error)
	Get(ctx context.Context, filter AccountFilter) (accs []Account, err error)
}

// AccountUseCase represents an account use case
type AccountUseCase interface {
	Create(ctx context.Context, acc Account) (id string, err error)
	Get(ctx context.Context, filter AccountFilter) (accs []Account, err error)
}

// TransactionRepository represents a transaction repository
type TransactionRepository interface {
	Create(ctx context.Context, tr Transaction) (err error)
	Get(ctx context.Context, tr TransactionFilter) ([]*Transaction, error)
}

// TransactionUseCase represents a transaction use case
type TransactionUseCase interface {
	Create(ctx context.Context, tr Transaction) (id string, err error)
	Get(ctx context.Context, tr TransactionFilter) (TransactionsPaging, error)
}
