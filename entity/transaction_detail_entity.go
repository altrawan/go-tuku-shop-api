package entity

import "gorm.io/gorm"

type TransactionDetail struct {
	ID            uint64 `gorm:"primary_key:auto_increment" json:"id"`
	TransactionID uint64 `gorm:"not null" json:"transaction_id"`
	ProductID     uint64 `gorm:"not null" json:"product_id"`
	Price         uint64 `gorm:"type:integer not null" json:"price"`
	Qty           uint64 `gorm:"type:integer not null" json:"qty"`
	gorm.Model
	Transaction Transaction `gorm:"foreignkey:TransactionID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"transaction"`
	Product     Product     `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product"`
}
