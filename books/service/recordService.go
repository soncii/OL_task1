package service

import (
	"context"
	"errors"
	"login/model"
	"login/storage"
	"time"
)

type RecordService struct {
	r *storage.Storage
}
type IRecordService interface {
	Get(ctx context.Context)

	CreateRecord(ctx context.Context, req model.RecordCreateReq) (*model.Record, error)
	Delete(ctx context.Context)
	UpdateRecord(ctx context.Context, record *model.Record) error
	GetBorrowedBooks(ctx context.Context) ([]model.BookWithRevenue, error)
}

func NewRecordService(r *storage.Storage) *RecordService {
	return &RecordService{r: r}
}
func (s *RecordService) Get(ctx context.Context) {
}

func (s *RecordService) CreateRecord(ctx context.Context, req model.RecordCreateReq) (*model.Record, error) {
	uid := req.UID
	bid := req.BID

	_, err := s.r.UserRepo.GetUserByID(ctx, uint(uid))
	if err != nil {
		return &model.Record{}, errors.New("user doesn't exist")
	}
	_, err = s.r.BookRepo.GetBookByID(ctx, uint(bid))
	if err != nil {
		return &model.Record{}, errors.New("book doesn't exist")
	}
	byBID, _ := s.r.RecordRepo.GetBorrowedRecordByBID(ctx, bid)
	if byBID.ID != 0 {
		return &model.Record{}, errors.New("the book is already borrowed")
	}
	r := model.Record{UserID: uid, BookID: bid, TakenAt: time.Now(), Borrowed: true}
	err = s.r.RecordRepo.CreateRecord(ctx, &r)
	if err != nil {
		return &model.Record{}, err
	}
	return &r, nil
}

func (*RecordService) Delete(ctx context.Context) {

}
func (s *RecordService) UpdateRecord(ctx context.Context, record *model.Record) error {
	return s.r.RecordRepo.UpdateRecord(ctx, record)
}
func (s *RecordService) GetBorrowedBooks(ctx context.Context) ([]model.BookWithRevenue, error) {
	return s.r.RecordRepo.GetBorrowedBooks(ctx)
}
