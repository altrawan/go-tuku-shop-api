package service

import (
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/repository"
)

type ProductService interface {
	List() []entity.Product
}

type iProductService struct {
	repository repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &iProductService{r}
}

func (s *iProductService) List() []entity.Product {
	return s.repository.List()
}
