package postgre

import (
	"context"
	"fmt"
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
	return r.db.Transaction(func(tx *gorm.DB) error {
		err := r.db.Create(book).Error
		if err != nil {
			return err
		}
		return r.db.Create(&model.BookRevenue{BookID: book.ID, Revenue: 0}).Error
	})
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
func (r *BookRepository) UpdateRevenue(ctx context.Context, revenue model.BookRevenue) error {
	fmt.Print("Printing Revenue: ")
	fmt.Println(revenue.Revenue)
	return r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&model.BookRevenue{}).Where("book_id = ?", revenue.BookID).Updates(map[string]interface{}{
			"revenue": gorm.Expr("revenue + ?", revenue.Revenue),
		}).Error
	})
}
