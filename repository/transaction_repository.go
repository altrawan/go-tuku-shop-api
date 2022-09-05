package repository

import (
	"go-tuku-shop-api/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	List() []entity.Transaction
	FindByPK(id uint64) entity.Transaction
	FindByUserID(UserID uint64) entity.Transaction
	Store(t entity.Transaction, td entity.TransactionDetail) (entity.Transaction, entity.TransactionDetail)
	UpdateAddress(t entity.Transaction) entity.Transaction
	UpdatePayment(t entity.Transaction) entity.Transaction
	Delete(b entity.Transaction)
}

type iTransactionRepository struct {
	connection *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &iTransactionRepository{db}
}

func (db *iTransactionRepository) List() []entity.Transaction {
	var Transactions []entity.Transaction
	db.connection.Preload("Transactions").Find(&Transactions)
	return Transactions
}

func (db *iTransactionRepository) FindByPK(id uint64) entity.Transaction {
	var Transaction entity.Transaction
	db.connection.Preload("Transactions").Find(&Transaction, id)
	return Transaction
}

func (db *iTransactionRepository) FindByUserID(UserID uint64) entity.Transaction {
	var transaction entity.Transaction
	db.connection.Where("user_id = ?", UserID).First(&transaction)
	return transaction
}

func (db *iTransactionRepository) Store(t entity.Transaction, td entity.TransactionDetail) (entity.Transaction, entity.TransactionDetail) {
	db.connection.Save(&t)
	db.connection.Preload("Transactions").Find(&t)

	td.TransactionID = t.ID
	db.connection.Save(&td)
	db.connection.Preload("Transaction_Details").Find(&td)
	return t, td
}

func (db *iTransactionRepository) UpdateAddress(t entity.Transaction) entity.Transaction {
	var transaction entity.Transaction
	transaction.RecipientName = t.RecipientName
	transaction.RecipientPhone = t.RecipientPhone
	transaction.City = t.City
	transaction.Address = t.Address
	transaction.PostalCode = t.PostalCode
	db.connection.Where("id = ?", t.ID).Model(&t).Updates(transaction)
	db.connection.Preload("Transactions").Find(&t)
	return t
}

func (db *iTransactionRepository) UpdatePayment(t entity.Transaction) entity.Transaction {
	var transaction entity.Transaction
	transaction.PaymentMethod = t.PaymentMethod
	db.connection.Where("id = ?", t.ID).Model(&t).Updates(transaction)
	db.connection.Preload("Transactions").Find(&t)
	return t
}

func (db *iTransactionRepository) Delete(b entity.Transaction) {
	db.connection.Delete(&b)
}
