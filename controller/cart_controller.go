package controller

import "gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/service"

type CartController interface {
}

type iCartController struct {
	service service.CartService
}

func NewCartController(s service.CartService) CartController {
	return &iCartController{s}
}
