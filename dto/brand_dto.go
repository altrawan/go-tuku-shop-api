package dto

// BrandCreateDTO is a model that user use when create a new category
type BrandCreateDTO struct {
	BrandName string `json:"brand_name" form:"brand_name" binding:"required"`
	Photo     string `json:"photo" form:"photo"`
}

// BrandUpdateDTO is a model that user use  when updating a category
type BrandUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	BrandName string `json:"brand_name" form:"brand_name" binding:"required"`
	Photo     string `json:"photo" form:"photo"`
}
