//go:generate mockgen -source=storage.go -destination=../mock_storage/mock_storage.go -package=mock_storage
package storage

import (
	"context"

	"gorm.io/gorm"

	"login/config"
	"login/model"
	"login/storage/postgre"
)

type IUserRepository interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUsersLastMonth(ctx context.Context, time string) ([]*model.User, error)
	GetUserByID(ctx context.Context, UID uint) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserRecords(ctx context.Context, email string) ([]*model.Record, error)
	Update(ctx context.Context, user *model.User) error
	CreateUser(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, uid uint, hard bool) error
}
type IBookRepository interface {
	GetBookByID(ctx context.Context, bid uint) (*model.Book, error)
	GetBookByTitle(ctx context.Context, title string) (*model.Book, error)
	CreateBook(ctx context.Context, book *model.Book) error
	DeleteBook(ctx context.Context, bid uint) error
	UpdateRevenue(ctx context.Context, revenue model.BookRevenue) error
}
type IRecordRepository interface {
	GetRecordsByEmail(ctx context.Context, email string) ([]*model.Record, error)
	CreateRecord(ctx context.Context, record *model.Record) error
	DeleteRecord(ctx context.Context) error
	UpdateRecord(ctx context.Context, record *model.Record) error
	GetBorrowedRecordByBID(ctx context.Context, bid uint) (*model.Record, error)
	GetBorrowedBooks(ctx context.Context) ([]model.BookWithRevenue, error)
	GetRecordByID(ctx context.Context, rid int) (*model.Record, error)
}

type Storage struct {
	DB         *gorm.DB
	UserRepo   IUserRepository
	BookRepo   IBookRepository
	RecordRepo IRecordRepository
}

func NewStorage(cfg *config.Config) (*Storage, error) {
	if cfg.DB == "pg" {
		db, err := postgre.Dial(cfg)
		db = db.Debug()
		if err != nil {
			return nil, err
		}
		return &Storage{DB: db, UserRepo: postgre.NewUserRepository(db), BookRepo: postgre.NewBookRepository(db),
			RecordRepo: postgre.NewRecordRepository(db)}, nil
	}
	return &Storage{}, nil
}
