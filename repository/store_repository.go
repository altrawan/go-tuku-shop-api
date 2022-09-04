package repository

import (
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/entity"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/helper"
	"gorm.io/gorm"
)

type StoreRepository interface {
	Update(u entity.User, s entity.Store) (entity.User, entity.Store)
}

type iStoreRepository struct {
	connection *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &iStoreRepository{db}
}

func (db *iStoreRepository) Update(u entity.User, s entity.Store) (entity.User, entity.Store) {
	if u.Password != "" {
		u.Password = helper.HashAndSalt([]byte(u.Password))
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, u.ID)
		u.Password = tempUser.Password
	}
	db.connection.Save(&u)

	s.UserID = u.ID
	db.connection.Save(&s)

	return u, s
}
