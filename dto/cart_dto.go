package dto

// CartCreateDTO is a model that user use when create a new category
type CartCreateDTO struct {
	UserID    uint64 `json:"user_id" form:"user_id"`
	ProductID uint64 `json:"product_id" form:"product_id" binding:"required"`
	Qty       uint64 `json:"qty" form:"qty" binding:"required"`
}

// CartUpdateDTO is a model that user use  when updating a category
type CartUpdateDTO struct {
	ID        uint64 `json:"id" form:"id"`
	UserID    uint64 `json:"user_id" form:"user_id"`
	ProductID uint64 `json:"product_id" form:"product_id" binding:"required"`
	Qty       uint64 `json:"qty" form:"qty" binding:"required"`
}
