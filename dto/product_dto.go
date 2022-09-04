package dto

// ProductCreateDTO is a model that user use when create a new category
type ProductCreateDTO struct {
	StoreID      uint64            `json:"store_id" form:"store_id" binding:"required"`
	CategoryID   uint64            `json:"category_id" form:"category_id" binding:"required"`
	ProductName  string            `json:"Product_name" form:"Product_name" binding:"required"`
	BrandID      uint64            `json:"brand_id" form:"brand_id" binding:"required"`
	Price        uint64            `json:"price" form:"price" binding:"required"`
	IsNew        uint64            `json:"is_new" form:"is_new" binding:"required"`
	Description  string            `json:"description" form:"description" binding:"required"`
	Stock        uint64            `json:"stock" form:"stock" binding:"required"`
	Rating       uint64            `json:"rating" form:"rating"`
	ProductImage []ProductImageDTO `json:"product_image" form:"product_image"`
	ProductColor []ProductColorDTO `json:"product_color" form:"product_color"`
	ProductSize  []ProductSizeDTO  `json:"product_size" form:"product_size"`
}

// ProductUpdateDTO is a model that user use  when updating a category
type ProductUpdateDTO struct {
	ID           uint64            `json:"id" form:"id"`
	StoreID      uint64            `json:"store_id" form:"store_id" binding:"required"`
	CategoryID   uint64            `json:"category_id" form:"category_id" binding:"required"`
	ProductName  string            `json:"Product_name" form:"Product_name" binding:"required"`
	BrandID      uint64            `json:"brand_id" form:"brand_id" binding:"required"`
	Price        uint64            `json:"price" form:"price" binding:"required"`
	IsNew        uint64            `json:"is_new" form:"is_new" binding:"required"`
	Description  string            `json:"description" form:"description" binding:"required"`
	Stock        uint64            `json:"stock" form:"stock" binding:"required"`
	Rating       uint64            `json:"rating" form:"rating"`
	ProductImage []ProductImageDTO `json:"product_image" form:"product_image"`
	ProductColor []ProductColorDTO `json:"product_color" form:"product_color"`
	ProductSize  []ProductSizeDTO  `json:"product_size" form:"product_size"`
}

type ProductImageDTO struct {
	Photo string `json:"photo" form:"photo"`
}

type ProductColorDTO struct {
	ColorName  string `json:"color_name" form:"color_name"`
	ColorValue string `json:"color_value" form:"color_value"`
}

type ProductSizeDTO struct {
	Size string `json:"size" form:"size"`
}
