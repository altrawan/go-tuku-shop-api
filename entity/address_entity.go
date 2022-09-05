package entity

import (
	"gorm.io/gorm"
)

type Address struct {
	ID             uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID         uint64 `gorm:"not null" json:"user_id"`
	Label          string `gorm:"type:varchar(255) not null" json:"label"`
	RecipientName  string `gorm:"type:varchar(255) not null" json:"recipient_name"`
	RecipientPhone string `gorm:"type:varchar(255) not null" json:"recipient_phone"`
	City           string `gorm:"type:varchar(255) not null" json:"city"`
	Address        string `gorm:"type:text not null" json:"address"`
	PostalCode     uint64 `gorm:"type:integer not null" json:"postal_code"`
	IsPrimary      uint64 `gorm:"type:integer not null" json:"is_primary"`
	gorm.Model
	User User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
