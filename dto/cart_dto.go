package dto

// CartCreateDTO is a model that user use when create a new category
type CartCreateDTO struct {
	ProductID string `json:"product_id" form:"product_id" binding:"required"`
	Qty       string `json:"qty" form:"qty" binding:"required"`
}

// CartUpdateDTO is a model that user use  when updating a category
type CartUpdateDTO struct {
	ID        uint64 `json:"id" form:"id"`
	ProductID string `json:"product_id" form:"product_id" binding:"required"`
	Qty       string `json:"qty" form:"qty" binding:"required"`
}
