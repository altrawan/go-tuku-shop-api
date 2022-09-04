package dto

// StoreUpdateDTO is used by user when PUT update Store
type StoreUpdateDTO struct {
	UserID           uint64 `json:"user_id" form:"user_id"`
	Name             string `json:"name" form:"name" binding:"required"`
	Email            string `json:"email" form:"email"`
	StoreName        string `json:"store_name" form:"store_name"`
	StorePhone       string `json:"store_phone" form:"store_phone"`
	StoreDescription string `json:"store_description" form:"store_descrition"`
	Photo            string `json:"photo" form:"photo"`
}

// ProfileChangePasswordDTO is used by user when PUT change password
type StoreChangePasswordDTO struct {
	UserID          uint64 `json:"user_id" form:"user_id"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}
