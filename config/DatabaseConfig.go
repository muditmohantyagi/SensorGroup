package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GoConnect() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	mySQLHost := os.Getenv("DB_HOST")
	mySQLUser := os.Getenv("DB_USER")
	mySQLPass := os.Getenv("DB_PASS")
	mySQLDBName := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Kolkata",
		mySQLHost,
		mySQLUser,
		mySQLPass,
		mySQLDBName,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to Gorm database" + err.Error())
	}
	//migration of table structure
	return db
}
