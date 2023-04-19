package service

import (
	"context"
	"transactions/model"
	"transactions/storage"
)

type TransactionService struct {
	r *storage.Storage
}
type ITransactionService interface {
	Get(ctx context.Context, u uint) (*model.Transaction, error)
	Create(ctx context.Context, req model.TransactionCreateReq) (model.TransactionCreateResp, error)
	Delete(ctx context.Context) error
}

func NewTransactionService(r *storage.Storage) *TransactionService {
	return &TransactionService{r: r}
}
func (s *TransactionService) Get(ctx context.Context, u uint) (*model.Transaction, error) {
	return s.r.TransactionRepo.Get(ctx, u)
}

func (s *TransactionService) Create(ctx context.Context, req model.TransactionCreateReq) (model.TransactionCreateResp, error) {

	tr := &model.Transaction{
		UserID: req.UID,
		BookID: req.BID,
		Price:  req.Price,
	}
	err := s.r.TransactionRepo.CreateTransaction(ctx, tr)
	if err != nil {
		return model.TransactionCreateResp{}, err
	}
	return model.TransactionCreateResp{TID: tr.ID}, nil
}

func (*TransactionService) Delete(ctx context.Context) error {
	return nil
}
