package entity

import "gorm.io/gorm"

type ProductColor struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ProductID  uint64 `gorm:"not null" json:"product_id"`
	ColorName  string `gorm:"type:varchar(255) not null" json:"color_name"`
	ColorValue string `gorm:"type:varchar(255) not null" json:"color_value"`
	gorm.Model
	Product Product `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product"`
}
