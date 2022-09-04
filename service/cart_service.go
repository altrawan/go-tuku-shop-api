package service

import "go-tuku-shop-api/repository"

type CartService interface {
}

type iCartService struct {
	repository repository.CartRepository
}

func NewCartService(r repository.CartRepository) CartService {
	return &iCartService{r}
}
