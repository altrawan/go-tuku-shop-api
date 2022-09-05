package service

import (
	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type TransactionService interface {
	List() []entity.Transaction
	FindByPK(id uint64) entity.Transaction
	FindByUserID(UserID uint64) entity.Transaction
	Store(b dto.TransactionCreateDTO) interface{}
	UpdateAddress(b dto.TransactionUpdateAddressDTO) entity.Transaction
	UpdatePayment(b dto.TransactionUpdatePaymentDTO) entity.Transaction
	Delete(b entity.Transaction)
	IsAllowedToEdit(id uint64, userID uint64) bool
}

type iTransactionService struct {
	repository repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) TransactionService {
	return &iTransactionService{r}
}

func (s *iTransactionService) List() []entity.Transaction {
	return s.repository.List()
}

func (s *iTransactionService) FindByPK(id uint64) entity.Transaction {
	return s.repository.FindByPK(id)
}

func (s *iTransactionService) FindByUserID(userID uint64) entity.Transaction {
	return s.repository.FindByUserID(userID)
}

func (s *iTransactionService) Store(t dto.TransactionCreateDTO) interface{} {
	transaction := entity.Transaction{}
	err := smapping.FillStruct(&transaction, smapping.MapFields(&t))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	transactionDetail := entity.TransactionDetail{}
	errDetail := smapping.FillStruct(&transactionDetail, smapping.MapFields(&t))
	if errDetail != nil {
		log.Fatalf("Failed map %v: ", errDetail)
	}

	dataTransaction, dataDetail := s.repository.Store(transaction, transactionDetail)

	res := map[string]uint64{
		"id":    dataDetail.TransactionID,
		"Price": dataDetail.Price,
		"Qty":   dataDetail.Qty,
		"Total": dataTransaction.Total,
	}

	return res
}

func (s *iTransactionService) UpdateAddress(b dto.TransactionUpdateAddressDTO) entity.Transaction {
	Transaction := entity.Transaction{}
	err := smapping.FillStruct(&Transaction, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.repository.UpdateAddress(Transaction)
	return res
}

func (s *iTransactionService) UpdatePayment(b dto.TransactionUpdatePaymentDTO) entity.Transaction {
	Transaction := entity.Transaction{}
	err := smapping.FillStruct(&Transaction, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.repository.UpdatePayment(Transaction)
	return res
}

func (s *iTransactionService) Delete(b entity.Transaction) {
	s.repository.Delete(b)
}

func (s *iTransactionService) IsAllowedToEdit(id uint64, userID uint64) bool {
	p := s.repository.FindByUserID(userID)
	return id == p.UserID
}
