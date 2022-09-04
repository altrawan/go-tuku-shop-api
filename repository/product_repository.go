package repository

import (
	"go-tuku-shop-api/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	List() []entity.Product
}

type iProductRepository struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &iProductRepository{db}
}

func (db *iProductRepository) List() []entity.Product {
	var products []entity.Product

	db.connection.Preload("Products").Find(&products)

	return products
}
