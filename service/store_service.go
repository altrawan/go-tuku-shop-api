package service

import (
	"log"

	"github.com/mashingan/smapping"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/dto"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/entity"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/repository"
)

type StoreService interface {
	Update(p dto.StoreUpdateDTO) interface{}
}

type iStoreService struct {
	repository repository.StoreRepository
}

func NewStoreService(r repository.StoreRepository) StoreService {
	return &iStoreService{r}
}

func (s *iStoreService) Update(store dto.StoreUpdateDTO) interface{} {
	entityUser := entity.User{}
	errUser := smapping.FillStruct(&entityUser, smapping.MapFields(&store))
	if errUser != nil {
		log.Fatalf("Failed map %v", errUser)
	}

	entityStore := entity.Store{}
	err := smapping.FillStruct(&entityStore, smapping.MapFields(&store))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	user, Store := s.repository.Update(entityUser, entityStore)

	res := map[string]string{
		"name":  Store.Name,
		"email": user.Email,
	}

	return res
}
