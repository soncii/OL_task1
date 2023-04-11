package service

import (
	"login/entities"
	"login/model"
	"login/storage"
)

type TransactionService struct {
	r *storage.Storage
}

func NewTransactionService(r *storage.Storage) *TransactionService {
	return &TransactionService{r: r}
}
func (s *TransactionService) Get() {
}

func (s *TransactionService) Create(req model.TransactionCreateReq) (model.TransactionCreateResp, error) {

	tr := &entities.Transaction{UserID: req.UID}
	trDetails := req.Books
	err := s.r.TransactionRepo.CreateTransaction(tr, trDetails)
	if err != nil {
		return model.TransactionCreateResp{}, err
	}
	return model.TransactionCreateResp{TID: tr.ID, CreatedAt: tr.CreatedAt}, nil
}

func (*TransactionService) Delete() {

}
