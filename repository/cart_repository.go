package repository

import "gorm.io/gorm"

type CartRepository interface {
}

type iCartRepository struct {
	connection *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &iCartRepository{db}
}
