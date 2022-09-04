package controller

import "gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/service"

type AddressController interface {
}

type iAddressController struct {
	service service.AddressService
}

func NewAddressController(s service.AddressService) AddressController {
	return &iAddressController{s}
}
