package dto

// ProfileUpdateDTO is used by user when PUT update profile
type ProfileUpdateDTO struct {
	UserID uint64 `json:"user_id" form:"user_id"`
	Name   string `json:"name" form:"name" binding:"required"`
	Email  string `json:"email" form:"email"`
	Phone  string `json:"phone" form:"phone"`
	Gender string `json:"gender" form:"gender"`
	Photo  string `json:"photo" form:"photo"`
}
