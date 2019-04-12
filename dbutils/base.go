package utils

import (
	"fmt"

	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error doring load environment")
	}
	dbName := os.Getenv("db_name")
	dbPass := os.Getenv("db_pass")
	dbUser := os.Getenv("db_user")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbURL := os.Getenv("database_url")
	dbURI := fmt.Sprintf(dbURL, dbUser, dbPass, dbHost, dbPort, dbName)
	localDB, err := gorm.Open("mysql", dbURI)
	db = localDB
	if err != nil {
		fmt.Println("error during create connection", err)
	}
	// db.AutoMigrate(&client.Client{}, &offer.Offer{}, &bid.Bid{})
}

func GetDB() *gorm.DB {
	return db
}
