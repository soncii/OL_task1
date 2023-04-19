package postgre

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"transactions/model"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}
func (r *TransactionRepository) Get(ctx context.Context, u uint) (*model.Transaction, error) {
	res := &model.Transaction{}
	err := r.db.First(res, u).Error
	return res, err
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, tr *model.Transaction) error {
	fmt.Println("Hey")
	return r.db.Create(tr).Error
}

func (r *TransactionRepository) Delete(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
