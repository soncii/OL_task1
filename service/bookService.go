package service

import (
	"fmt"
	"login/entities"
	"login/model"
	"login/storage"
)

type BookService struct {
	r *storage.Storage
}

func NewBookService(r *storage.Storage) *BookService {
	return &BookService{r: r}
}
func (s *BookService) Get() {
}

func (s *BookService) Create(req model.BookCreateReq) model.BookCreateResp {
	b := entities.Book{Title: req.Title, Author: req.Author}
	fmt.Printf("Printing from service:%v\n", b)
	s.r.BookRepo.CreateBook(&b)
	return model.BookCreateResp{BID: b.ID, CreatedAt: b.CreatedAt}
}

func (*BookService) Delete() {

}
