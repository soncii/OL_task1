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

func (r *RecordRepository) UpdateRecord(ctx context.Context, record *model.Record) error {
	return r.db.Save(record).Error
}
func (r *RecordRepository) DeleteRecord(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
func (r *RecordRepository) GetBorrowedRecordByBID(ctx context.Context, bid uint) (*model.Record, error) {
	record := &model.Record{}
	return record, r.db.Where("book_id=? AND borrowed=true", bid).First(record).Error
}
func (r *RecordRepository) GetBorrowedBooks(ctx context.Context) ([]model.BookWithRevenue, error) {
	var Result []model.BookWithRevenue
	r.db.Raw("select b.id, b.title, b.author, br.revenue from records r " +
		"join books b on r.book_id=b.id " +
		"join book_revenues br on br.book_id=b.id where r.borrowed=true").Scan(&Result)
	return Result, nil
}
