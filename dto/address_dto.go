package dto

type AddressCreateDTO struct {
	UserID         uint64 `json:"user_id" form:"user_id"`
	Label          string `json:"label" form:"label" binding:"required"`
	RecipientName  string `json:"recipient_name" form:"recipient_name" binding:"required"`
	RecipientPhone string `json:"recipient_phone" form:"recipient_phone" binding:"required"`
	City           string `json:"city" form:"city" binding:"required"`
	Address        string `json:"address" form:"address" binding:"required"`
	PostalCode     uint64 `json:"postal_code" form:"postal_code" binding:"required"`
	IsPrimary      uint64 `json:"is_primary" form:"is_primary" binding:"required"`
}

// AddressUpdateDTO is a model that user use  when updating a category
type AddressUpdateDTO struct {
	ID             uint64 `json:"id" form:"id"`
	UserID         uint64 `json:"user_id" form:"user_id"`
	Label          string `json:"label" form:"label" binding:"required"`
	RecipientName  string `json:"recipient_name" form:"recipient_name" binding:"required"`
	RecipientPhone string `json:"recipient_phone" form:"recipient_phone" binding:"required"`
	City           string `json:"city" form:"city" binding:"required"`
	Address        string `json:"address" form:"address" binding:"required"`
	PostalCode     uint64 `json:"postal_code" form:"postal_code" binding:"required"`
	IsPrimary      uint64 `json:"is_primary" form:"is_primary" binding:"required"`
}
