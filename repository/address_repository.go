package repository

import (
	"go-tuku-shop-api/entity"

	"gorm.io/gorm"
)

type AddressRepository interface {
	List() []entity.Address
	FindByPK(id uint64) entity.Address
	FindByUserID(UserID uint64) entity.Address
	Store(b entity.Address) entity.Address
	Update(b entity.Address) entity.Address
	Delete(b entity.Address)
}

type iAddressRepository struct {
	connection *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &iAddressRepository{db}
}

func (db *iAddressRepository) List() []entity.Address {
	var Addresss []entity.Address
	db.connection.Preload("Addresses").Find(&Addresss)
	return Addresss
}

func (db *iAddressRepository) FindByPK(id uint64) entity.Address {
	var Address entity.Address
	db.connection.Preload("Addresses").Find(&Address, id)
	return Address
}

func (db *iAddressRepository) FindByUserID(UserID uint64) entity.Address {
	var address entity.Address
	db.connection.Where("user_id = ?", UserID).First(&address)
	return address
}

func (db *iAddressRepository) Store(b entity.Address) entity.Address {
	db.connection.Save(&b)
	db.connection.Preload("Addresses").Find(&b)
	return b
}

func (db *iAddressRepository) Update(b entity.Address) entity.Address {
	db.connection.Save(&b)
	db.connection.Preload("Addresses").Find(&b)
	return b
}

func (db *iAddressRepository) Delete(b entity.Address) {
	db.connection.Delete(&b)
}
