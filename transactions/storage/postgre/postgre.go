package postgre

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"transactions/config"
	"transactions/model"
)

func Dial(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort, cfg.DbSSL)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	err = db.AutoMigrate(&model.Transaction{})
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return db, nil
}
