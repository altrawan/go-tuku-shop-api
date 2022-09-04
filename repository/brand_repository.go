package repository

import (
	"go-tuku-shop-api/entity"

	"gorm.io/gorm"
)

type BrandRepository interface {
	List() []entity.Brand
	FindByPK(id uint64) entity.Brand
	Store(b entity.Brand) entity.Brand
	Update(b entity.Brand) entity.Brand
	Delete(b entity.Brand)
}

type iBrandRepository struct {
	connection *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &iBrandRepository{db}
}

func (db *iBrandRepository) List() []entity.Brand {
	var brands []entity.Brand
	db.connection.Preload("Brands").Find(&brands)
	return brands
}

func (db *iBrandRepository) FindByPK(id uint64) entity.Brand {
	var brand entity.Brand
	db.connection.Preload("Brands").Find(&brand, id)
	return brand
}

func (db *iBrandRepository) Store(b entity.Brand) entity.Brand {
	db.connection.Save(&b)
	db.connection.Preload("Brands").Find(&b)
	return b
}

func (db *iBrandRepository) Update(b entity.Brand) entity.Brand {
	db.connection.Save(&b)
	db.connection.Preload("Brands").Find(&b)
	return b
}

func (db *iBrandRepository) Delete(b entity.Brand) {
	db.connection.Delete(&b)
}
