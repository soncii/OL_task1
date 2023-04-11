package postgre

import (
	"gorm.io/gorm"
	"login/entities"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}
func (r *TransactionRepository) Get() error {
	//TODO implement me
	panic("implement me")
}

func (r *TransactionRepository) CreateTransaction(tr *entities.Transaction, trDetails []*entities.TransactionDetails) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(tr).Error; err != nil {
			return err
		}
		for _, detail := range trDetails {
			detail.ID = tr.ID
			if err := tx.Create(detail).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}
