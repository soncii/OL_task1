package postgre

import (
	"context"
	"gorm.io/gorm"
	"login/model"
)

type RecordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) *RecordRepository {
	return &RecordRepository{db: db}
}
func (r *RecordRepository) GetRecordsByEmail(ctx context.Context, email string) ([]*model.Record, error) {
	record := make([]*model.Record, 0)
	err := r.db.Where("email=?", email).Find(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (r *RecordRepository) CreateRecord(ctx context.Context, record *model.Record) error {
	return r.db.Create(record).Error
}

func (r *RecordRepository) DeleteRecord(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
