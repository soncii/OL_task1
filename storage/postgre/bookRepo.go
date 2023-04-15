package postgre

import (
	"context"
	"gorm.io/gorm"
	"login/model"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetBookByTitle(ctx context.Context, title string) (*model.Book, error) {
	b := &model.Book{}
	err := r.db.Where("title=?", title).First(b).Error
	if err != nil {
		return &model.Book{}, err
	}
	return b, nil
}

func (r *BookRepository) CreateBook(ctx context.Context, book *model.Book) error {
	return r.db.Create(book).Error

}

func (r *BookRepository) DeleteBook(ctx context.Context, bid uint) error {
	err := r.db.Unscoped().Delete(&model.Book{}, bid).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *BookRepository) GetBookByID(ctx context.Context, bid uint) (*model.Book, error) {
	b := &model.Book{}
	err := r.db.Where("id=?", bid).First(b).Error
	if err != nil {
		return &model.Book{}, err
	}
	return b, nil
}
