package postgre

import (
	"gorm.io/gorm"
	"login/entities"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetBookByTitle(title string) *entities.Book {
	b := &entities.Book{}
	r.db.Where("title=?", title).First(b)
	return b
}

func (r *BookRepository) CreateBook(book *entities.Book) {
	r.db.Create(book)

}

func (r *BookRepository) DeleteBook() {
	//TODO implement me
	panic("implement me")
}
