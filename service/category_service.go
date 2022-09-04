package service

import (
	"log"

	"github.com/mashingan/smapping"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/dto"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/entity"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/repository"
)

type CategoryService interface {
	List() []entity.Category
	Store(b dto.CategoryCreateDTO) entity.Category
	Update(b dto.CategoryUpdateDTO) entity.Category
	Delete(b entity.Category)
	FindByID(id uint64) entity.Category
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

func (s *iCategoryService) FindByID(id uint64) entity.Category {
	return s.repository.FindByID(id)
}
