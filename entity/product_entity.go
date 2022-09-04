package entity

import (
	"database/sql"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint64        `gorm:"primary_key:auto_increment" json:"id"`
	StoreID     uint64        `gorm:"not null" json:"store_id"`
	CategoryID  uint64        `gorm:"not null" json:"category_id"`
	ProductName string        `gorm:"type:varchar(255) not null" json:"product_name"`
	BrandID     uint64        `gorm:"not null" json:"brand_id"`
	Price       uint64        `gorm:"type:integer not null" json:"price"`
	IsNew       uint64        `gorm:"type:enum('0','1') not null" json:"is_new"`
	Description string        `gorm:"type:text not null" json:"description"`
	Stock       uint64        `gorm:"type:integer not null" json:"stock"`
	Rating      sql.NullInt64 `gorm:"type:integer" json:"rating"`
	gorm.Model
	Store    Store    `gorm:"foreignkey:StoreID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"store"`
	Ctaegory Category `gorm:"foreignkey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category"`
	Brand    Brand    `gorm:"foreignkey:BrandID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"brand"`
}
