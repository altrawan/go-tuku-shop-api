package dto

// CategoryCreateDTO is a model that user use when create a new category
type CategoryCreateDTO struct {
	CategoryName string `json:"Category_name" form:"Category_name" binding:"required"`
	Photo        string `json:"photo" form:"photo"`
}

// CategoryUpdateDTO is a model that user use  when updating a category
type CategoryUpdateDTO struct {
	ID           uint64 `json:"id" form:"id" binding:"required"`
	CategoryName string `json:"Category_name" form:"Category_name" binding:"required"`
	Photo        string `json:"photo" form:"photo"`
}
