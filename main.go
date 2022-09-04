package main

import (
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/config"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/router"
)

func main() {
	db := config.NewClient()

	sqlDB, _ := db.DB()

	defer sqlDB.Close()

	r := router.NewRouter(db)

	r.Run()
}
