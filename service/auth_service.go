package service

import (
	"log"

	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"
	"go-tuku-shop-api/repository"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	Login(email string, password string) interface{}
	RegisterSeller(r dto.RegisterSellerDTO) interface{}
	RegisterBuyer(r dto.RegisterBuyerDTO) interface{}
	IsDuplicateEmail(email string) bool
}

type iAuthService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) AuthService {
	return &iAuthService{r}
}

func (s *iAuthService) Login(email string, password string) interface{} {
	res := s.repository.Login(email, password)
	if u, ok := res.(entity.User); ok {
		comparedPassword := helper.VerifyPassword(u.Password, []byte(password))

		if u.Email == email && comparedPassword {
			return res
		}

		return false
	}

	return false
}

func (s *iAuthService) RegisterSeller(r dto.RegisterSellerDTO) interface{} {
	entityUser := entity.User{}
	errUser := smapping.FillStruct(&entityUser, smapping.MapFields(&r))
	if errUser != nil {
		log.Fatalf("Failed map %v", errUser)
	}

	entityStore := entity.Store{}
	err := smapping.FillStruct(&entityStore, smapping.MapFields(&r))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	user, store := s.repository.RegisterSeller(entityUser, entityStore)

	res := map[string]string{
		"name":        store.Name,
		"email":       user.Email,
		"store_name":  store.StoreName,
		"store_phone": store.StorePhone,
	}

	return res
}

func (s *iAuthService) RegisterBuyer(r dto.RegisterBuyerDTO) interface{} {
	entityUser := entity.User{}
	errUser := smapping.FillStruct(&entityUser, smapping.MapFields(&r))
	if errUser != nil {
		log.Fatalf("Failed map %v", errUser)
	}

	entityProfile := entity.Profile{}
	err := smapping.FillStruct(&entityProfile, smapping.MapFields(&r))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	user, profile := s.repository.RegisterBuyer(entityUser, entityProfile)

	res := map[string]string{
		"name":  profile.Name,
		"email": user.Email,
	}

	return res
}

func (s *iAuthService) IsDuplicateEmail(email string) bool {
	res := s.repository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
