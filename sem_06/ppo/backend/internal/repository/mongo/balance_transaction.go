package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/mongo/models"
	"github.com/google/uuid"
)

type BalanceTransactionRepository struct {
	db *mongo.Database
}

func NewBTR(db *mongo.Database) BalanceTransactionRepository {
	return BalanceTransactionRepository{db: db}
}

func (b BalanceTransactionRepository) Create(ctx context.Context, transaction *models.BalanceTransaction) error {
	coll := b.db.Collection("balance_transaction")

	tran := &repoModels.BalanceTransaction{
		UUID:   uuid.New(),
		Reason: transaction.Reason,
		Time:   transaction.Time,
		Amount: transaction.Amount,
		UserID: transaction.UserID,
	}

	_, err := coll.InsertOne(ctx, tran)
	if err != nil {
		return fmt.Errorf("insert: %w", err)
	}

	return nil
}
