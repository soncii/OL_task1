package storage

import (
	"context"
	//_ "github.com/golang/mock/mockgen/model"
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
	Delete(ctx context.Context) error
}
type IBookRepository interface {
	GetBookByTitle(ctx context.Context, title string) (*model.Book, error)
	CreateBook(ctx context.Context, book *model.Book) error
	DeleteBook(ctx context.Context) error
}
type IRecordRepository interface {
	GetRecordsByEmail(ctx context.Context, email string) ([]*model.Record, error)
	CreateRecord(ctx context.Context, record *model.Record) error
	DeleteRecord(ctx context.Context) error
}

type Storage struct {
	DB         *gorm.DB
	UserRepo   IUserRepository
	BookRepo   IBookRepository
	RecordRepo IRecordRepository
}

func NewStorage(cfg *config.Config) *Storage {
	if cfg.DB == "pg" {
		db, err := postgre.Dial(cfg)
		db = db.Debug()
		if err != nil {
			return nil
		}
		return &Storage{DB: db, UserRepo: postgre.NewUserRepository(db), BookRepo: postgre.NewBookRepository(db),
			RecordRepo: postgre.NewRecordRepository(db)}
	}
	return nil
}
