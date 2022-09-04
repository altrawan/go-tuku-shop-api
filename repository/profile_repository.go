package repository

import (
	"time"

	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/entity"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Update(u entity.User, p entity.Profile) (entity.User, entity.Profile)
	FindByID(profileID uint64) entity.Profile
}

type iProfileRepository struct {
	connection *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &iProfileRepository{db}
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

func (db *iProfileRepository) FindByID(profileID uint64) entity.Profile {
	var profile entity.Profile
	db.connection.Preload("Profiles").Find(&profile, profileID)
	return profile
}
