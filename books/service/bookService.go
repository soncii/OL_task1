package service

import (
	"context"
	"login/model"
	"login/storage"
)

type IBookService interface {
	GetByID(ctx context.Context, bid model.BookGetByIDReq) (*model.Book, error)
	Create(ctx context.Context, req model.BookCreateReq) (model.BookCreateResp, error)
	DeleteByID(ctx context.Context, bid model.BookGetByIDReq) (bool, error)
	UpdateBookRevenue(ctx context.Context, revenue model.BookRevenue) error
}

type BookService struct {
	r *storage.Storage
}

func NewBookService(r *storage.Storage) *BookService {
	return &BookService{r: r}
}
func (s *BookService) GetByID(ctx context.Context, bid model.BookGetByIDReq) (*model.Book, error) {
	book, err := s.r.BookRepo.GetBookByID(ctx, bid.BID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookService) Create(ctx context.Context, req model.BookCreateReq) (model.BookCreateResp, error) {
	b := model.Book{Title: req.Title, Author: req.Author}
	err := s.r.BookRepo.CreateBook(ctx, &b)
	if err != nil {
		return model.BookCreateResp{}, err
	}
	return model.BookCreateResp{BID: b.ID, CreatedAt: b.CreatedAt}, nil
}

func (s *BookService) DeleteByID(ctx context.Context, bid model.BookGetByIDReq) (bool, error) {
	err := s.r.BookRepo.DeleteBook(ctx, bid.BID)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (s *BookService) UpdateBookRevenue(ctx context.Context, revenue model.BookRevenue) error {
	return s.r.BookRepo.UpdateRevenue(ctx, revenue)
}
