package entity

import (
	"gorm.io/gorm"
)

type Profile struct {
	ID     uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID uint64 `gorm:"not null" json:"user_id"`
	Name   string `gorm:"type:varchar(255) not null" json:"name"`
	Phone  string `gorm:"type:varchar(255)" json:"phone"`
	Gender string `gorm:"type:enum('male','female')" json:"gender"`
	Photo  string `gorm:"type:varchar(255)" json:"photo"`
	gorm.Model
	User User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
