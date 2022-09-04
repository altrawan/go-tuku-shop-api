package entity

import "gorm.io/gorm"

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Email    string `gorm:"type:varchar(255) not null" json:"email"`
	Password string `gorm:"type:varchar(255) not null" json:"password"`
	Level    string `gorm:"type:enum('admin','seller','buyer')" json:"level"`
	gorm.Model
}
