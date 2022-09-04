package entity

import "gorm.io/gorm"

type Category struct {
	ID           uint64 `gorm:"primary_key:auto_increment" json:"id"`
	CategoryName string `gorm:"type:varchar(255) not null" json:"category_name"`
	Photo        string `gorm:"type:varchar(255)" json:"photo"`
	gorm.Model
}
