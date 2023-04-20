package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"login/model"
	"login/storage"
)

type RecordService struct {
	r *storage.Storage
}
type IRecordService interface {
	Get(ctx context.Context, rid int, url string) (*model.RecordWithTransaction, error)
	CreateRecord(ctx context.Context, req model.RecordCreateReq, url string) (*model.Record, error)
	Delete(ctx context.Context)
	UpdateRecord(ctx context.Context, record *model.Record) error
	GetBorrowedBooks(ctx context.Context) ([]model.BookWithRevenue, error)
}

func NewRecordService(r *storage.Storage) *RecordService {
	return &RecordService{r: r}
}
func (s *RecordService) Get(ctx context.Context, rid int, url string) (*model.RecordWithTransaction, error) {
	record, err := s.r.RecordRepo.GetRecordByID(ctx, rid)
	if err != nil {
		return nil, err
	}
	r, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/transaction/%s", url, strconv.Itoa(int(record.TransactionID))), nil)
	r.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	var respBody model.TransactionGetResp
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}
	return record.MapWithPrice(respBody.Price), nil
}

func (s *RecordService) CreateRecord(ctx context.Context, req model.RecordCreateReq, url string) (*model.Record, error) {
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
	r := &model.Record{UserID: uid, BookID: bid, TakenAt: time.Now(), Borrowed: true}
	err = s.r.RecordRepo.CreateRecord(ctx, r)
	if err != nil {
		return &model.Record{}, err
	}
	if err != nil {
		return nil, err
	}
	reqBody := []byte(fmt.Sprintf(`{"UID":%d, "BID":%v, "Price":%v}`,
		req.UID, req.BID, req.Price))
	reqTr, _ := http.NewRequest(http.MethodPost, url+"/transaction", bytes.NewBuffer(reqBody))
	reqTr.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(reqTr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var createResp model.TransactionCreateResp
	if err := json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		return nil, err
	}
	r.TransactionID = createResp.TID
	err = s.r.RecordRepo.UpdateRecord(ctx, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (*RecordService) Delete(ctx context.Context) {

}
func (s *RecordService) UpdateRecord(ctx context.Context, record *model.Record) error {
	return s.r.RecordRepo.UpdateRecord(ctx, record)
}
func (s *RecordService) GetBorrowedBooks(ctx context.Context) ([]model.BookWithRevenue, error) {
	return s.r.RecordRepo.GetBorrowedBooks(ctx)
}
