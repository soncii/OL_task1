package storage

import (
	"context"
	"gorm.io/gorm"
	"transactions/config"
	"transactions/model"
	"transactions/storage/postgre"
)

type ITransactionRepository interface {
	Get(ctx context.Context, u uint) (*model.Transaction, error)
	CreateTransaction(ctx context.Context, tr *model.Transaction) error
	Delete(ctx context.Context) error
}
type Storage struct {
	DB              *gorm.DB
	TransactionRepo ITransactionRepository
}

func NewStorage(cfg *config.Config) *Storage {
	if cfg.DB == "pg" {
		db, err := postgre.Dial(cfg)
		db = db.Debug()
		if err != nil {
			return nil
		}
		return &Storage{DB: db, TransactionRepo: postgre.NewTransactionRepository(db)}
	}
	return nil
}
