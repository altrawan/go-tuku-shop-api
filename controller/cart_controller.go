package controller

import "go-tuku-shop-api/service"

type CartController interface {
}

type iCartController struct {
	service service.CartService
}

func NewCartController(s service.CartService) CartController {
	return &iCartController{s}
}
