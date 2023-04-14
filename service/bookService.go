package service

import (
	"context"
	"errors"
	"fmt"
	"login/model"
	"login/storage"
)

type IBookService interface {
	Get(ctx context.Context) error
	Create(ctx context.Context, req model.BookCreateReq) (model.BookCreateResp, error)
	Delete(ctx context.Context) error
}

type BookService struct {
	r *storage.Storage
}

func NewBookService(r *storage.Storage) *BookService {
	return &BookService{r: r}
}
func (s *BookService) Get(ctx context.Context) error {
	return errors.New("IMPLEMENT ME!")
}

func (s *BookService) Create(ctx context.Context, req model.BookCreateReq) (model.BookCreateResp, error) {
	b := model.Book{Title: req.Title, Author: req.Author}
	fmt.Printf("Printing from service:%v\n", b)
	err := s.r.BookRepo.CreateBook(ctx, &b)
	if err != nil {
		return model.BookCreateResp{}, err
	}
	return model.BookCreateResp{BID: b.ID, CreatedAt: b.CreatedAt}, nil
}

func (*BookService) Delete(ctx context.Context) error {
	return errors.New("IMPLEMENT ME!")
}
