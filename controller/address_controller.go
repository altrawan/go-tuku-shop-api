package controller

import "go-tuku-shop-api/service"

type AddressController interface {
}

type iAddressController struct {
	service service.AddressService
}

func NewAddressController(s service.AddressService) AddressController {
	return &iAddressController{s}
}
