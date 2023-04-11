package service

import (
	"login/entities"
	"login/model"
	"login/storage"
	"strconv"
	"time"
)

type RecordService struct {
	r *storage.Storage
}

func NewRecordService(r *storage.Storage) *RecordService {
	return &RecordService{r: r}
}
func (s *RecordService) Get() {
}

func (s *RecordService) Create(req model.RecordCreateReq) (model.RecordCreateResp, error) {
	uid, err := strconv.Atoi(req.UID)
	if err != nil {
		return model.RecordCreateResp{}, err
	}
	bid, err1 := strconv.Atoi(req.BID)
	if err1 != nil {
		return model.RecordCreateResp{}, err
	}
	r := entities.Record{UserID: uint(uid), BookID: uint(bid), TakenAt: time.Now(), Borrowed: true}
	s.r.RecordRepo.CreateRecord(&r)
	return model.RecordCreateResp{RID: r.ID, CreatedAt: r.CreatedAt}, nil
}

func (*RecordService) Delete() {

}
