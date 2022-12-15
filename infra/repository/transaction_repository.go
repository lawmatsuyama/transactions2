package repository

import (
	"context"
	"time"

	"github.com/lawmatsuyama/pismo-transactions/domain"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	batchSizeTransaction int32 = 20
)

// TransactionRepository implements domain.TransactionRepository interface. It contains a database client.
type TransactionRepository struct {
	Client *mongo.Client
}

// NewTransactionRepository returns a new TransactionRepository
func NewTransactionRepository(client *mongo.Client) TransactionRepository {
	return TransactionRepository{
		Client: client,
	}
}

// Create receive domain.Transaction and insert it in database collection
func (db TransactionRepository) Create(ctx context.Context, tr domain.Transaction) error {
	l := log.WithField("transaction", tr)
	c := db.Client.Database("account").Collection("transactions")
	_, err := c.InsertOne(ctx, tr)
	if err != nil {
		l.WithError(err).Error("failed to insert transaction")
		return err
	}

	return nil
}

// Get receives a filter of transaction, query it in database and returns the result in domain.Transaction
func (db TransactionRepository) Get(ctx context.Context, filterTr domain.TransactionFilter) (trs []*domain.Transaction, err error) {
	l := log.WithField("filter", filterTr)
	c := db.Client.Database("account").Collection("transactions")
	filter := bson.D{}
	filter = filterSimple(filter, "_id", filterTr.ID, isZeroComparable[string])
	filter = filterSimple(filter, "account_id", filterTr.AccountID, isZeroComparable[string])
	filter = filterSimple(filter, "operation_type_id", filterTr.OperationTypeID, isZeroComparable[domain.OperationType])
	filter = filterRange(filter, "amount", filterTr.AmountGreater, filterTr.AmountLess, isZeroComparable[float64])
	filter = filterSimple(filter, "description", filterTr.Description, isZeroComparable[string])
	filter = filterRange(filter, "event_date", filterTr.EventDateFrom, filterTr.EventDateTo, isZeroTime)

	sort := bson.D{bsonE("created_at", 1), bsonE("_id", 1)}
	opts := options.Find().
		SetSort(sort).
		SetBatchSize(batchSizeTransaction).
		SetMaxTime(time.Second * 20).
		SetSkip(filterTr.Paging.Skip()).
		SetLimit(filterTr.Paging.LimitByPage())

	cur, err := c.Find(ctx, filter, opts)
	if err == mongo.ErrNoDocuments {
		err = domain.ErrTransactionsNotFound
		return
	}

	if err != nil {
		l.WithError(err).Error("Failed to get transactions")
		err = domain.ErrUnknown
		return
	}

	trs = []*domain.Transaction{}
	err = cur.All(ctx, &trs)
	if err != nil {
		l.WithError(err).Error("Failed to iterate over transactions")
		err = domain.ErrUnknown
		return
	}

	return trs, err

}
