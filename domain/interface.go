package domain

type TransactionValidator interface {
	IsValid() error
}

// UUIDGenerator represents an UUID generator
type UUIDGenerator interface {
	Generate() string
}
