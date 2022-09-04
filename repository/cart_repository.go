package repository

import (
	"go-tuku-shop-api/entity"

	"gorm.io/gorm"
)

type CartRepository interface {
	List() []entity.Cart
	FindByPK(id uint64) entity.Cart
	Store(b entity.Cart) entity.Cart
	Update(b entity.Cart) entity.Cart
	Delete(b entity.Cart)
}

type iCartRepository struct {
	connection *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &iCartRepository{db}
}

func (db *iCartRepository) List() []entity.Cart {
	var Carts []entity.Cart
	db.connection.Preload("Carts").Find(&Carts)
	return Carts
}

func (db *iCartRepository) FindByPK(id uint64) entity.Cart {
	var Cart entity.Cart
	db.connection.Preload("Carts").Find(&Cart, id)
	return Cart
}

func (db *iCartRepository) Store(b entity.Cart) entity.Cart {
	db.connection.Save(&b)
	db.connection.Preload("Carts").Find(&b)
	return b
}

func (db *iCartRepository) Update(b entity.Cart) entity.Cart {
	db.connection.Save(&b)
	db.connection.Preload("Carts").Find(&b)
	return b
}

func (db *iCartRepository) Delete(b entity.Cart) {
	db.connection.Delete(&b)
}
