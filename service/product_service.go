package service

import (
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/entity"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/repository"
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
