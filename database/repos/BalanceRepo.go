package repos

import (
	"errors"

	"gorm.io/gorm"
)

type Balance struct {
	ID     int64 `gorm:"primaryKey;column:id"`
	Amount int64 `gorm:"column:amount"`
	UserID int64 `gorm:"column:user_id"`
}

func (Balance) TableName() string {
	return "balance"
}

type BalanceRepo struct {
	DB *gorm.DB
}

func (r *BalanceRepo) CreateBalance(userID int64) (*Balance, error) {
	balance := &Balance{
		UserID: userID,
		Amount: 0,
	}
	result := r.DB.Create(balance)
	return balance, result.Error
}

func (r *BalanceRepo) AddBalance(userID int64, amount int64) (*Balance, error) {
	var balance Balance
	result := r.DB.Where("user_id = ?", userID).First(&balance)
	if result.Error != nil {
		return nil, errors.New("balance not found")
	}

	balance.Amount += amount
	result = r.DB.Model(&balance).Update("amount", balance.Amount)
	return &balance, result.Error
}

func (r *BalanceRepo) SubtractBalance(userID int64, amount int64) (*Balance, error) {
	var balance Balance
	result := r.DB.Where("user_id = ?", userID).First(&balance)
	if result.Error != nil {
		return nil, errors.New("balance not found")
	}

	if balance.Amount < amount {
		return nil, errors.New("insufficient funds")
	}

	balance.Amount -= amount
	result = r.DB.Model(&balance).Update("amount", balance.Amount)
	return &balance, result.Error
}
