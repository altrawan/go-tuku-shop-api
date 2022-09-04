package entity

import (
	"database/sql"

	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint64         `gorm:"primary_key:auto_increment" json:"id"`
	UserID         uint64         `gorm:"not null" json:"user_id"`
	Invoice        string         `gorm:"type:varchar(255) not null" json:"invoice"`
	Total          uint64         `gorm:"type:integer not null" json:"total"`
	PaymentMethod  sql.NullString `gorm:"type:varchar(255)" json:"payment_method"`
	Status         string         `gorm:"type:varchar(255) not null" json:"status"`
	RecipientName  sql.NullString `gorm:"type:varchar(255)" json:"recipient_name"`
	RecipientPhone sql.NullString `gorm:"type:varchar(255)" json:"recipient_phone"`
	City           sql.NullString `gorm:"type:varchar(255)" json:"city"`
	Address        sql.NullString `gorm:"type:text" json:"address"`
	PostalCode     sql.NullInt64  `gorm:"type:integer" json:"postal_code"`
	gorm.Model
	User User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
