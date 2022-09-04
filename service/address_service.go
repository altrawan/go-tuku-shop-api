package service

import "go-tuku-shop-api/repository"

type AddressService interface {
}

type iAddressService struct {
	repository repository.AddressRepository
}

func NewAddressService(r repository.AddressRepository) AddressService {
	return &iAddressService{r}
}
