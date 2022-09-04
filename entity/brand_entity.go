package entity

import (
	"gorm.io/gorm"
)

type Brand struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	BrandName string `gorm:"type:varchar(255) not null" json:"brand_name"`
	Photo     string `gorm:"type:varchar(255)" json:"photo"`
	gorm.Model
}
