package dto

// CategoryCreateDTO is a model that user use when create a new category
type CategoryCreateDTO struct {
	CategoryName string `json:"category_name" form:"category_name" binding:"required"`
	Photo        string `json:"photo" form:"photo"`
}

// CategoryUpdateDTO is a model that user use  when updating a category
type CategoryUpdateDTO struct {
	ID           uint64 `json:"id" form:"id"`
	CategoryName string `json:"category_name" form:"category_name" binding:"required"`
	Photo        string `json:"photo" form:"photo"`
}
