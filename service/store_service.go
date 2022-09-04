package service

import (
	"log"

	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/repository"

	"github.com/mashingan/smapping"
)

type StoreService interface {
	List() []entity.Store
	FindByPK(StoreID uint64) entity.Store
	FindByUserID(StoreID uint64) entity.Store
	Update(d dto.StoreUpdateDTO) interface{}
	ChangePassword(d dto.StoreChangePasswordDTO) interface{}
	IsAllowedToEdit(userID uint64, StoreID uint64) bool
}

type iStoreService struct {
	repository repository.StoreRepository
}

func NewStoreService(r repository.StoreRepository) StoreService {
	return &iStoreService{r}
}

func (s *iStoreService) List() []entity.Store {
	return s.repository.List()
}

func (s *iStoreService) FindByPK(StoreID uint64) entity.Store {
	return s.repository.FindByPK(StoreID)
}

func (s *iStoreService) FindByUserID(UserID uint64) entity.Store {
	return s.repository.FindByUserID(UserID)
}

func (s *iStoreService) Update(d dto.StoreUpdateDTO) interface{} {
	entityUser := entity.User{}
	errUser := smapping.FillStruct(&entityUser, smapping.MapFields(&d))
	if errUser != nil {
		log.Fatalf("Failed map %v", errUser)
	}

	entityStore := entity.Store{}
	err := smapping.FillStruct(&entityStore, smapping.MapFields(&d))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	user, Store := s.repository.Update(entityUser, entityStore)

	res := map[string]string{
		"name":              Store.Name,
		"email":             user.Email,
		"store_name":        Store.StoreName,
		"store_phone":       Store.StorePhone,
		"store_description": Store.StoreDescription,
	}

	return res
}

func (s *iStoreService) ChangePassword(d dto.StoreChangePasswordDTO) interface{} {
	entityUser := entity.User{}
	errUser := smapping.FillStruct(&entityUser, smapping.MapFields(&d))
	if errUser != nil {
		log.Fatalf("Failed map %v", errUser)
	}

	entityStore := entity.Store{}
	err := smapping.FillStruct(&entityStore, smapping.MapFields(&d))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	user, _ := s.repository.ChangePassword(entityUser, entityStore)

	res := map[string]string{
		"password": user.Password,
	}

	return res
}

func (s *iStoreService) IsAllowedToEdit(userID uint64, StoreID uint64) bool {
	p := s.repository.FindByUserID(StoreID)
	return userID == p.UserID
}
