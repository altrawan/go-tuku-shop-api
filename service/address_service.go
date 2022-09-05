package service

import (
	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type AddressService interface {
	List() []entity.Address
	FindByPK(id uint64) entity.Address
	FindByUserID(UserID uint64) entity.Address
	Store(b dto.AddressCreateDTO) entity.Address
	Update(b dto.AddressUpdateDTO) entity.Address
	Delete(b entity.Address)
	IsAllowedToEdit(id uint64, userID uint64) bool
}

type iAddressService struct {
	repository repository.AddressRepository
}

func NewAddressService(r repository.AddressRepository) AddressService {
	return &iAddressService{r}
}

func (s *iAddressService) List() []entity.Address {
	return s.repository.List()
}

func (s *iAddressService) FindByPK(id uint64) entity.Address {
	return s.repository.FindByPK(id)
}

func (s *iAddressService) FindByUserID(userID uint64) entity.Address {
	return s.repository.FindByUserID(userID)
}

func (s *iAddressService) Store(b dto.AddressCreateDTO) entity.Address {
	Address := entity.Address{}
	err := smapping.FillStruct(&Address, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.repository.Store(Address)
	return res
}

func (s *iAddressService) Update(b dto.AddressUpdateDTO) entity.Address {
	Address := entity.Address{}
	err := smapping.FillStruct(&Address, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.repository.Update(Address)
	return res
}

func (s *iAddressService) Delete(b entity.Address) {
	s.repository.Delete(b)
}

func (s *iAddressService) IsAllowedToEdit(id uint64, userID uint64) bool {
	p := s.repository.FindByUserID(userID)
	return id == p.UserID
}
