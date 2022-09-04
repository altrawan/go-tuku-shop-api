package entity

import (
	"gorm.io/gorm"
)

type ProductImage struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ProductID uint64 `gorm:"not null" json:"product_id"`
	Photo     string `gorm:"type:varchar(255)" json:"photo"`
	gorm.Model
	Product Product `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product"`
}
