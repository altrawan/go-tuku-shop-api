package repository

import (
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(email string, password string) interface{}
	RegisterSeller(u entity.User, s entity.Store) (entity.User, entity.Store)
	RegisterBuyer(u entity.User, p entity.Profile) (entity.User, entity.Profile)
	IsDuplicateEmail(email string) (tx *gorm.DB)
}

type iAuthRepository struct {
	connection *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &iAuthRepository{db}
}

func (db *iAuthRepository) Login(email string, password string) interface{} {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *iAuthRepository) RegisterSeller(u entity.User, s entity.Store) (entity.User, entity.Store) {
	u.Password = helper.HashAndSalt([]byte(u.Password))
	u.Level = "seller"
	db.connection.Save(&u)

	s.UserID = u.ID
	db.connection.Save(&s)
	return u, s
}

func (db *iAuthRepository) RegisterBuyer(u entity.User, p entity.Profile) (entity.User, entity.Profile) {
	u.Password = helper.HashAndSalt([]byte(u.Password))
	u.Level = "buyer"
	db.connection.Save(&u)

	p.UserID = u.ID
	db.connection.Save(&p)
	return u, p
}

func (db *iAuthRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}
