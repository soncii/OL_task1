package postgre

import (
	"gorm.io/gorm"
	"login/entities"
)

type RecordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) *RecordRepository {
	return &RecordRepository{db: db}
}
func (r *RecordRepository) GetRecordsByEmail(email string) []*entities.Record {
	result := make([]*entities.Record, 10)
	return result
}

func (r *RecordRepository) CreateRecord(record *entities.Record) {
	r.db.Create(record)
}

func (r *RecordRepository) DeleteRecord() {
	//TODO implement me
	panic("implement me")
}
