package service

import "gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/repository"

type AddressService interface {
}

type iAddressService struct {
	repository repository.AddressRepository
}

func NewAddressService(r repository.AddressRepository) AddressService {
	return &iAddressService{r}
}
