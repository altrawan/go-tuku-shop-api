package repository

import (
	"go-tuku-shop-api/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	List() []entity.Category
	FindByPK(id uint64) entity.Category
	Store(c entity.Category) entity.Category
	Update(c entity.Category) entity.Category
	Delete(c entity.Category)
}

type iCategoryRepository struct {
	connection *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &iCategoryRepository{db}
}

func (db *iCategoryRepository) List() []entity.Category {
	var categories []entity.Category
	db.connection.Preload("Categories").Find(&categories)
	return categories
}

func (db *iCategoryRepository) FindByPK(id uint64) entity.Category {
	var Category entity.Category
	db.connection.Preload("Categories").Find(&Category, id)
	return Category
}

func (db *iCategoryRepository) Store(c entity.Category) entity.Category {
	db.connection.Save(&c)
	db.connection.Preload("Categories").Find(&c)
	return c
}

func (db *iCategoryRepository) Update(c entity.Category) entity.Category {
	db.connection.Save(&c)
	db.connection.Preload("Categories").Find(&c)
	return c
}

func (db *iCategoryRepository) Delete(c entity.Category) {
	db.connection.Delete(&c)
}
