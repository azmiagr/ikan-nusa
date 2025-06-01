package repository

import (
	"ikan-nusa/entity"

	"gorm.io/gorm"
)

type ITransactionRepository interface {
	CreateTransaction(tx *gorm.DB, transaction *entity.Transaction) (*entity.Transaction, error)
	CreateTransactionItems(tx *gorm.DB, transactionItems *entity.TransactionItems) (*entity.TransactionItems, error)
}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{db: db}
}

func (t *TransactionRepository) CreateTransaction(tx *gorm.DB, transaction *entity.Transaction) (*entity.Transaction, error) {
	err := tx.Debug().Create(&transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionRepository) CreateTransactionItems(tx *gorm.DB, transactionItems *entity.TransactionItems) (*entity.TransactionItems, error) {
	err := tx.Debug().Create(&transactionItems).Error
	if err != nil {
		return nil, err
	}

	return transactionItems, nil
}
