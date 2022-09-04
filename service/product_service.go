package service

import (
	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type ProductService interface {
	List() []entity.Product
	FindByPK(id uint64) entity.Product
	Store(b dto.ProductCreateDTO) entity.Product
	Update(b dto.ProductUpdateDTO) entity.Product
	Delete(b entity.Product)
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

func (s *iProductService) FindByPK(id uint64) entity.Product {
	return s.repository.FindByPK(id)
}

func (s *iProductService) Store(b dto.ProductCreateDTO) entity.Product {
	product := entity.Product{}
	errProduct := smapping.FillStruct(&product, smapping.MapFields(&b))
	if errProduct != nil {
		log.Fatalf("Failed map %v: ", errProduct)
	}

	res := s.repository.Store(product)

	for _, val := range b.ProductImage {
		productImage := entity.ProductImage{
			ProductID: res.ID,
			Photo:     val.Photo,
		}
		s.repository.StoreImage(productImage)
	}

	for _, val := range b.ProductColor {
		productColor := entity.ProductColor{
			ProductID:  res.ID,
			ColorName:  val.ColorName,
			ColorValue: val.ColorValue,
		}
		s.repository.StoreColor(productColor)
	}

	for _, val := range b.ProductSize {
		productSize := entity.ProductSize{
			ProductID: res.ID,
			Size:      val.Size,
		}
		s.repository.StoreSize(productSize)
	}

	return res
}

func (s *iProductService) Update(b dto.ProductUpdateDTO) entity.Product {
	Product := entity.Product{}
	err := smapping.FillStruct(&Product, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	res := s.repository.Update(Product)

	s.repository.DeleteImage(res.ID)
	for _, val := range b.ProductImage {
		productImage := entity.ProductImage{
			ProductID: res.ID,
			Photo:     val.Photo,
		}
		s.repository.StoreImage(productImage)
	}

	s.repository.DeleteColor(res.ID)
	for _, val := range b.ProductColor {
		productColor := entity.ProductColor{
			ProductID:  res.ID,
			ColorName:  val.ColorName,
			ColorValue: val.ColorValue,
		}
		s.repository.StoreColor(productColor)
	}

	s.repository.DeleteSize(res.ID)
	for _, val := range b.ProductSize {
		productSize := entity.ProductSize{
			ProductID: res.ID,
			Size:      val.Size,
		}
		s.repository.StoreSize(productSize)
	}

	return res
}

func (s *iProductService) Delete(b entity.Product) {
	s.repository.Delete(b)
	s.repository.DeleteImage(b.ID)
	s.repository.DeleteColor(b.ID)
	s.repository.DeleteSize(b.ID)
}
