package migration

import (
	"go-tuku-shop-api/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{}, &entity.Profile{}, &entity.Store{}, &entity.Address{}, &entity.Chat{}, &entity.Category{}, &entity.Brand{}, &entity.Product{}, &entity.ProductSize{}, &entity.ProductImage{}, &entity.ProductColor{}, &entity.Cart{}, &entity.Transaction{}, &entity.TransactionDetail{})
}
