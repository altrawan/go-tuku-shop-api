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

// ProfileChangePasswordDTO is used by user when PUT change password
type ProfileChangePasswordDTO struct {
	UserID          uint64 `json:"user_id" form:"user_id"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}
