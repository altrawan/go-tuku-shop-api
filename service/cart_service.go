package service

import (
	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type CartService interface {
	List() []entity.Cart
	FindByPK(id uint64) entity.Cart
	FindByUserID(UserID uint64) entity.Cart
	Store(b dto.CartCreateDTO) entity.Cart
	Update(b dto.CartUpdateDTO) entity.Cart
	Delete(b entity.Cart)
	IsAllowedToEdit(id uint64, userID uint64) bool
}

type iCartService struct {
	repository repository.CartRepository
}

func NewCartService(r repository.CartRepository) CartService {
	return &iCartService{r}
}

func (s *iCartService) List() []entity.Cart {
	return s.repository.List()
}

func (s *iCartService) FindByPK(id uint64) entity.Cart {
	return s.repository.FindByPK(id)
}

func (s *iCartService) FindByUserID(userID uint64) entity.Cart {
	return s.repository.FindByUserID(userID)
}

func (s *iCartService) Store(b dto.CartCreateDTO) entity.Cart {
	Cart := entity.Cart{}
	err := smapping.FillStruct(&Cart, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.repository.Store(Cart)
	return res
}

func (s *iCartService) Update(b dto.CartUpdateDTO) entity.Cart {
	Cart := entity.Cart{}
	err := smapping.FillStruct(&Cart, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.repository.Update(Cart)
	return res
}

func (s *iCartService) Delete(b entity.Cart) {
	s.repository.Delete(b)
}

func (s *iCartService) IsAllowedToEdit(id uint64, userID uint64) bool {
	p := s.repository.FindByUserID(userID)
	return id == p.UserID
}
