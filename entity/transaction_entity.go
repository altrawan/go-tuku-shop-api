package entity

import (
	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID         uint64 `gorm:"not null" json:"user_id"`
	Invoice        string `gorm:"type:varchar(255) not null" json:"invoice"`
	Total          uint64 `gorm:"type:integer not null" json:"total"`
	PaymentMethod  string `gorm:"type:varchar(255)" json:"payment_method"`
	Status         string `gorm:"type:varchar(255) not null" json:"status"`
	RecipientName  string `gorm:"type:varchar(255)" json:"recipient_name"`
	RecipientPhone string `gorm:"type:varchar(255)" json:"recipient_phone"`
	City           string `gorm:"type:varchar(255)" json:"city"`
	Address        string `gorm:"type:text" json:"address"`
	PostalCode     uint64 `gorm:"type:integer" json:"postal_code"`
	gorm.Model
	User User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
