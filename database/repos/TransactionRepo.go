package repos

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	TransactionID int64     `gorm:"primaryKey;column:transaction_id"`
	TrTime        time.Time `gorm:"column:tr_time"`
	Amount        int64     `gorm:"column:amount"`
	BalanceID     int64     `gorm:"column:balance_id"`
}

func (Transaction) TableName() string {
	return "transactions"
}

type TransactionRepo struct {
	DB *gorm.DB
}

func (r *TransactionRepo) CreateLog(balanceID int64, amount int64) (*Transaction, error) {
	transaction := &Transaction{
		BalanceID: balanceID,
		Amount:    amount,
		TrTime:    time.Now(),
	}
	result := r.DB.Create(transaction)
	return transaction, result.Error
}

func (r *TransactionRepo) GetLogsByBalanceID(balanceID int64) ([]Transaction, error) {
	var transactions []Transaction
	result := r.DB.Where("balance_id = ?", balanceID).Find(&transactions)
	return transactions, result.Error
}

func (r *TransactionRepo) GetLogsByUserID(userID string) ([]Transaction, error) {
	var balance Balance
	r.DB.Where("user_id = ?", userID).Find(&balance)

	transactions, result := r.GetLogsByBalanceID(balance.ID)
	return transactions, result
}
