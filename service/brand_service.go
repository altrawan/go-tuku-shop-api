package service

import (
	"log"

	"github.com/mashingan/smapping"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/dto"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/entity"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/repository"
)

type BrandService interface {
	List() []entity.Brand
	Store(b dto.BrandCreateDTO) entity.Brand
	Update(b dto.BrandUpdateDTO) entity.Brand
	Delete(b entity.Brand)
	FindByID(id uint64) entity.Brand
}

type iBrandService struct {
	repository repository.BrandRepository
}

func NewBrandService(r repository.BrandRepository) BrandService {
	return &iBrandService{r}
}

func (s *iBrandService) List() []entity.Brand {
	return s.repository.List()
}

func (s *iBrandService) Store(b dto.BrandCreateDTO) entity.Brand {
	brand := entity.Brand{}
	err := smapping.FillStruct(&brand, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.repository.Store(brand)
	return res
}

func (s *iBrandService) Update(b dto.BrandUpdateDTO) entity.Brand {
	brand := entity.Brand{}
	err := smapping.FillStruct(&brand, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.repository.Update(brand)
	return res
}

func (s *iBrandService) Delete(b entity.Brand) {
	s.repository.Delete(b)
}

func (s *iBrandService) FindByID(id uint64) entity.Brand {
	return s.repository.FindByID(id)
}
