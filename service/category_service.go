package service

import (
	"log"

	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/repository"

	"github.com/mashingan/smapping"
)

type CategoryService interface {
	List() []entity.Category
	FindByPK(id uint64) entity.Category
	Store(b dto.CategoryCreateDTO) entity.Category
	Update(b dto.CategoryUpdateDTO) entity.Category
	Delete(b entity.Category)
}

type iCategoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &iCategoryService{r}
}

func (s *iCategoryService) List() []entity.Category {
	return s.repository.List()
}

func (s *iCategoryService) FindByPK(id uint64) entity.Category {
	return s.repository.FindByPK(id)
}

func (s *iCategoryService) Store(b dto.CategoryCreateDTO) entity.Category {
	category := entity.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.repository.Store(category)
	return res
}

func (s *iCategoryService) Update(b dto.CategoryUpdateDTO) entity.Category {
	category := entity.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.repository.Update(category)
	return res
}

func (s *iCategoryService) Delete(b entity.Category) {
	s.repository.Delete(b)
}
