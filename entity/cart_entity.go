package entity

import "gorm.io/gorm"

type Cart struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID    uint64 `gorm:"not null" json:"user_id"`
	ProductID uint64 `gorm:"not null" json:"product_id"`
	Qty       uint64 `gorm:"type:integer not null" json:"qty"`
	gorm.Model
	User    User    `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Product Product `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product"`
}
