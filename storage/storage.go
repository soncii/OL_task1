package storage

import (
	"gorm.io/gorm"
	"login/config"
	"login/entities"
	"login/storage/postgre"
)

type IUserRepository interface {
	GetUsers() []*entities.User
	GetUsersLastMonth(time string) []*entities.User
	GetUserByID(UID uint) *entities.User
	GetUserByEmail(email string) *entities.User
	GetUserRecords(email string) []*entities.Record
	Update(user *entities.User) error
	CreateUser(user *entities.User)
	Delete()
}
type IBookRepository interface {
	GetBookByTitle(title string) *entities.Book
	CreateBook(book *entities.Book)
	DeleteBook()
}
type IRecordRepository interface {
	GetRecordsByEmail(email string) []*entities.Record
	CreateRecord(record *entities.Record)
	DeleteRecord()
}
type ITransactionRepository interface {
	Get() error
	CreateTransaction(tr *entities.Transaction, trDetails []*entities.TransactionDetails) error
	Delete() error
}
type Storage struct {
	DB              *gorm.DB
	UserRepo        IUserRepository
	BookRepo        IBookRepository
	RecordRepo      IRecordRepository
	TransactionRepo ITransactionRepository
}

func NewStorage(cfg *config.Config) *Storage {
	if cfg.DB == "pg" {
		db, err := postgre.Dial(cfg)
		db = db.Debug()
		if err != nil {
			return nil
		}
		return &Storage{DB: db, UserRepo: postgre.NewUserRepository(db), BookRepo: postgre.NewBookRepository(db),
			RecordRepo: postgre.NewRecordRepository(db), TransactionRepo: postgre.NewTransactionRepository(db)}
	}
	return nil
}
