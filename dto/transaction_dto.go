package dto

// TransactionCreateDTO is a model that user use when create a new transaction
type TransactionCreateDTO struct {
	UserID        uint64 `json:"user_id" form:"user_id"`
	Invoice       string `json:"invoice" form:"invoice"`
	Total         uint64 `json:"total" form:"total"`
	Status        string `json:"status" form:"status"`
	TransactionID uint64 `json:"transaction_id" form:"transaction_id"`
	ProductID     uint64 `json:"product_id" form:"product_id"`
	Price         uint64 `json:"price" form:"price"`
	Qty           uint64 `json:"qty" form:"qty"`
}

// TransactionUpdateAddressDTO is a model that user use  when updating a address transaction
type TransactionUpdateAddressDTO struct {
	ID             uint64 `json:"id" form:"id"`
	RecipientName  string `json:"recipient_name" form:"recipient_name" binding:"required"`
	RecipientPhone string `json:"recipient_phone" form:"recipient_phone" binding:"required"`
	City           string `json:"city" form:"city" binding:"required"`
	Address        string `json:"address" form:"address" binding:"required"`
	PostalCode     uint64 `json:"postal_code" form:"postal_code"`
}

// TransactionUpdatePaymentDTO is a model that user use  when updating a payment transaction
type TransactionUpdatePaymentDTO struct {
	ID            uint64 `json:"id" form:"id"`
	PaymentMethod string `json:"payment_method" form:"payment_method" binding:"required"`
}
