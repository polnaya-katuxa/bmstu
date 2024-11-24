package postgres

import (
	"context"
	"fmt"

	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BalanceTransactionRepository struct {
	db *gorm.DB
}

func NewBTR(db *gorm.DB) BalanceTransactionRepository {
	return BalanceTransactionRepository{db: db}
}

func (b BalanceTransactionRepository) Create(ctx context.Context, transaction *models.BalanceTransaction) error {
	tran := &repoModels.BalanceTransaction{
		UUID:   uuid.New(),
		Reason: transaction.Reason,
		Time:   transaction.Time,
		Amount: transaction.Amount,
		UserID: transaction.UserID,
	}

	res := b.db.WithContext(ctx).Table("balance_transactions").Create(tran)
	if res.Error != nil {
		return toErrorResponse(fmt.Errorf("insert: %w", res.Error))
	}

	return nil
}
