package config

import (
	"fmt"
	"log"
	"os"

	"go-tuku-shop-api/migration"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewClient is creating a new connection to our database
func NewClient() *gorm.DB {
	if os.Getenv("ENVIRONMENT") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if os.Getenv("ENVIRONMENT") == "production" {
		dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=require"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		migration.Migrate(db)

		return db
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	migration.Migrate(db)

	return db
}

// CloseClient method is closing a connection between application and database
func CloseClient(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
