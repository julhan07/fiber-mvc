package infra

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func posgresCreateConnection(dbHost, dbPort, dbName, dbUser, dbPassword string) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// defer db.DB(

	config, err := db.DB()
	if err != nil {
		panic(err)
	}

	config.SetMaxOpenConns(5)
	config.SetMaxIdleConns(5)
	config.SetConnMaxIdleTime(20 * time.Second)
	config.SetConnMaxLifetime(5 * time.Minute)

	// defer config.Close()

	return db
}

func GetDB() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db := posgresCreateConnection(dbHost, dbPort, dbName, dbUser, dbPassword)

	return db
}
