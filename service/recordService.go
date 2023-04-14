package service

import (
	"context"
	"login/model"
	"login/storage"
	"strconv"
	"time"
)

type RecordService struct {
	r *storage.Storage
}
type IRecordService interface {
	Get(ctx context.Context)
	Create(ctx context.Context, req model.RecordCreateReq) (model.RecordCreateResp, error)
	Delete(ctx context.Context)
}

func NewRecordService(r *storage.Storage) *RecordService {
	return &RecordService{r: r}
}
func (s *RecordService) Get(ctx context.Context) {
}

func (s *RecordService) Create(ctx context.Context, req model.RecordCreateReq) (model.RecordCreateResp, error) {
	uid, err := strconv.Atoi(req.UID)
	if err != nil {
		return model.RecordCreateResp{}, err
	}
	bid, err1 := strconv.Atoi(req.BID)
	if err1 != nil {
		return model.RecordCreateResp{}, err
	}
	r := model.Record{UserID: uint(uid), BookID: uint(bid), TakenAt: time.Now(), Borrowed: true}
	err = s.r.RecordRepo.CreateRecord(ctx, &r)
	if err != nil {
		return model.RecordCreateResp{}, err
	}
	return model.RecordCreateResp{RID: r.ID, CreatedAt: r.CreatedAt}, nil
}

func (*RecordService) Delete(ctx context.Context) {

}
