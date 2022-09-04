package service

import "gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/repository"

type CartService interface {
}

type iCartService struct {
	repository repository.CartRepository
}

func NewCartService(r repository.CartRepository) CartService {
	return &iCartService{r}
}
