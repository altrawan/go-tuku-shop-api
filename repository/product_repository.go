package repository

import (
	"go-tuku-shop-api/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	List() []entity.Product
	FindByPK(id uint64) entity.Product
	Store(p entity.Product) entity.Product
	StoreImage(p entity.ProductImage)
	StoreColor(p entity.ProductColor)
	StoreSize(p entity.ProductSize)
	Update(b entity.Product) entity.Product
	Delete(b entity.Product)
	DeleteImage(productID uint64)
	DeleteColor(productID uint64)
	DeleteSize(productID uint64)
}

type iProductRepository struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &iProductRepository{db}
}

func (db *iProductRepository) List() []entity.Product {
	var Products []entity.Product
	db.connection.Preload("Products").Find(&Products)
	return Products
}

func (db *iProductRepository) FindByPK(id uint64) entity.Product {
	var Product entity.Product
	db.connection.Preload("Products").Find(&Product, id)
	return Product
}

func (db *iProductRepository) Store(p entity.Product) entity.Product {
	db.connection.Save(&p)
	db.connection.Preload("Products").Find(&p)

	return p
}

func (db *iProductRepository) StoreImage(p entity.ProductImage) {
	db.connection.Save(&p)
	db.connection.Preload("Product_Images").Find(&p)
}

func (db *iProductRepository) StoreColor(p entity.ProductColor) {
	db.connection.Save(&p)
	db.connection.Preload("Product_Colors").Find(&p)
}

func (db *iProductRepository) StoreSize(p entity.ProductSize) {
	db.connection.Save(&p)
	db.connection.Preload("Product_Sizes").Find(&p)
}

func (db *iProductRepository) Update(b entity.Product) entity.Product {
	db.connection.Save(&b)
	db.connection.Preload("Products").Find(&b)
	return b
}

func (db *iProductRepository) Delete(b entity.Product) {
	db.connection.Delete(&b)
}

func (db *iProductRepository) DeleteImage(productID uint64) {
	db.connection.Where("product_id = ?", productID).Delete(&entity.ProductImage{})
}

func (db *iProductRepository) DeleteColor(productID uint64) {
	db.connection.Where("product_id = ?", productID).Delete(&entity.ProductColor{})
}

func (db *iProductRepository) DeleteSize(productID uint64) {
	db.connection.Where("product_id = ?", productID).Delete(&entity.ProductSize{})
}
