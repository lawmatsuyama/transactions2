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
	batchSizeAccount int32 = 20
)

// AccountRepository implements domain.AccountRepository interface. It contains the database client.
type AccountRepository struct {
	Client *mongo.Client
}

// NewAccountRepository returns a new AccountRepository
func NewAccountRepository(client *mongo.Client) AccountRepository {
	return AccountRepository{
		Client: client,
	}
}

// Create receive domain.Account and insert it in database collection
func (db AccountRepository) Create(ctx context.Context, acc domain.Account) error {
	l := log.WithField("accounts", acc)
	c := db.Client.Database("account").Collection("accounts")
	_, err := c.InsertOne(ctx, acc)
	if err != nil {
		l.WithError(err).Error("failed to insert account")
		return err
	}

	return nil
}

// Get receives a filter of account, query it in database and returns the result in domain.Account
func (db AccountRepository) Get(ctx context.Context, filterAcc domain.AccountFilter) (accs []domain.Account, err error) {
	l := log.WithField("filter", filterAcc)
	c := db.Client.Database("account").Collection("accounts")
	filter := bson.D{}
	filter = filterSimple(filter, "_id", filterAcc.ID, isZeroComparable[string])
	filter = filterSimple(filter, "document_number", filterAcc.DocumentNumber, isZeroComparable[domain.DocumentNumber])
	filter = filterRange(filter, "created_at", filterAcc.CreatedAtFrom, filterAcc.CreatedAtTo, isZeroTime)

	sort := bson.D{bsonE("created_at", 1), bsonE("_id", 1)}
	opts := options.Find().
		SetSort(sort).
		SetBatchSize(batchSizeAccount).
		SetMaxTime(time.Second * 20).
		SetSkip(filterAcc.Paging.Skip()).
		SetLimit(filterAcc.Paging.LimitByPage())

	cur, err := c.Find(ctx, filter, opts)
	if err == mongo.ErrNoDocuments {
		err = domain.ErrAccountNotFound
		return
	}

	if err != nil {
		l.WithError(err).Error("Failed to get accounts")
		err = domain.ErrUnknown
		return
	}

	accs = []domain.Account{}
	err = cur.All(ctx, &accs)
	if err != nil {
		l.WithError(err).Error("Failed to iterate over accounts")
		err = domain.ErrUnknown
		return
	}

	return accs, err

}
