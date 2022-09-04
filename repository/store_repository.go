package repository

import (
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"
	"time"

	"gorm.io/gorm"
)

type StoreRepository interface {
	List() []entity.Store
	FindByPK(StoreID uint64) entity.Store
	FindByUserID(UserID uint64) entity.Store
	Update(u entity.User, s entity.Store) (entity.User, entity.Store)
	ChangePassword(u entity.User, p entity.Store) (entity.User, entity.Store)
}

type iStoreRepository struct {
	connection *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &iStoreRepository{db}
}

func (db *iStoreRepository) List() []entity.Store {
	var Stores []entity.Store
	db.connection.Preload("Stores").Find(&Stores)
	return Stores
}

func (db *iStoreRepository) FindByPK(StoreID uint64) entity.Store {
	var store entity.Store
	db.connection.Preload("Stores").Find(&store, StoreID)
	return store
}

func (db *iStoreRepository) FindByUserID(UserID uint64) entity.Store {
	var store entity.Store
	db.connection.Where("user_id = ?", UserID).First(&store)
	return store
}

func (db *iStoreRepository) Update(u entity.User, s entity.Store) (entity.User, entity.Store) {
	var user entity.User
	user.Email = u.Email
	user.UpdatedAt = time.Now()
	db.connection.Where("id = ?", s.UserID).Model(&u).Updates(user)

	var Store entity.Store
	Store.Name = s.Name
	Store.StoreName = s.StoreName
	Store.StorePhone = s.StorePhone
	Store.StoreDescription = s.StoreDescription
	db.connection.Where("user_id = ?", s.UserID).Model(&s).Updates(Store)

	return u, s
}

func (db *iStoreRepository) ChangePassword(u entity.User, s entity.Store) (entity.User, entity.Store) {
	var user entity.User
	user.Password = helper.HashAndSalt([]byte(u.Password))
	db.connection.Where("id = ?", s.UserID).Model(&u).Updates(user)
	return u, s
}
