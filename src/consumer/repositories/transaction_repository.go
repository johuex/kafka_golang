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

func (t *TransactionRepository) InsertTransactions(transactions []TransactionDB) *gorm.DB {
	result := t.DBConn.Create(&transactions)
	t.DBConn.Commit()
	return result
}
