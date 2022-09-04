package entity

import "gorm.io/gorm"

type Chat struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Sender   uint64 `gorm:"not null" json:"sender"`
	Receiver uint64 `gorm:"not null" json:"receiver"`
	Message  string `gorm:"type:text not null" json:"message"`
	gorm.Model
	UserSender   User `gorm:"foreignkey:Sender;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user_sender"`
	UserReceiver User `gorm:"foreignkey:Receiver;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user_receiver"`
}
