package repositories

import (
	"time"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DBConn *gorm.DB
}

func NewTransactionRepository(dbConn *gorm.DB) *TransactionRepository {
	return &TransactionRepository{dbConn}
}

type TransactionDB struct {
	gorm.Model
	TransactionType string
	Location        string
	TransactionId   int64
	AccountNumber   int64
	Amount          int
	Time            time.Time
	Accepted        bool
}

func (t *TransactionRepository) InsertTransactions(transactions []TransactionDB) error {
	tx := t.DBConn.Begin()
	if err := tx.Create(&transactions).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil

}
