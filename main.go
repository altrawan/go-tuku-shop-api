package main

import (
	"go-tuku-shop-api/config"
	"go-tuku-shop-api/router"
)

func main() {
	db := config.NewClient()

	sqlDB, _ := db.DB()

	defer sqlDB.Close()

	r := router.NewRouter(db)

	r.Run()
}
