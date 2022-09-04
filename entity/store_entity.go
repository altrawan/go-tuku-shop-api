package entity

import (
	"gorm.io/gorm"
)

type Store struct {
	ID               uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID           uint64 `gorm:"not null" json:"user_id"`
	Name             string `gorm:"type:varchar(255) not null" json:"name"`
	StoreName        string `gorm:"type:varchar(255) not null" json:"store_name"`
	StorePhone       string `gorm:"type:varchar(255) not null" json:"store_phone"`
	StoreDescription string `gorm:"type:text" json:"store_description"`
	gorm.Model
	User User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
