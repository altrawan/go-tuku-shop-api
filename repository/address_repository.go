package repository

import "gorm.io/gorm"

type AddressRepository interface {
}

type iAddressRepository struct {
	connection *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &iAddressRepository{db}
}
