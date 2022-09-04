package dto

// LoginDTO is a model that used by client when POST from /login url
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"passsword" binding:"required"`
}

// RegisterSellerDTO is used when client post /register-seller url
type RegisterSellerDTO struct {
	Name       string `json:"name" form:"name" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	StoreName  string `json:"store_name" form:"store_name" binding:"required"`
	StorePhone string `json:"store_phone" form:"stire_phone" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
}

// RegisterBuyerDTO is used when client post from /register-buyer url
type RegisterBuyerDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
