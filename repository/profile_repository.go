package repository

import (
	"time"

	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	List() []entity.Profile
	FindByID(profileID uint64) entity.Profile
	Update(u entity.User, p entity.Profile) (entity.User, entity.Profile)
	ChangePassword(u entity.User) entity.User
}

type iProfileRepository struct {
	connection *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &iProfileRepository{db}
}

func (db *iProfileRepository) List() []entity.Profile {
	var profiles []entity.Profile
	db.connection.Preload("Profiles").Find(&profiles)
	return profiles
}

func (db *iProfileRepository) FindByID(profileID uint64) entity.Profile {
	var profile entity.Profile
	db.connection.Preload("Profiles").Find(&profile, profileID)
	return profile
}

func (db *iProfileRepository) Update(u entity.User, p entity.Profile) (entity.User, entity.Profile) {
	var user entity.User
	user.Email = u.Email
	user.UpdatedAt = time.Now()
	db.connection.Where("id = ?", p.UserID).Model(&u).Updates(user)

	var profile entity.Profile
	profile.Name = p.Name
	profile.Phone = p.Phone
	profile.Gender = p.Gender
	profile.Photo = p.Photo
	db.connection.Where("user_id = ?", p.UserID).Model(&p).Updates(profile)

	return u, p
}

func (db *iProfileRepository) ChangePassword(u entity.User) entity.User {
	var user entity.User
	user.Password = helper.HashAndSalt([]byte(u.Password))
	db.connection.Where("id = ?", u.ID).Model(&u).Updates(user)
	return u
}
