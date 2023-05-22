package config

import (
	"auth-service/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB_Instance struct {
	DB *gorm.DB
}

var PostgresDB DB_Instance

func ConnectDB() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_POSTGRES_USER"),
		os.Getenv("DB_POSTGRES_PASSWORD"),
		os.Getenv("DB_POSTGRES_NAME"),
	)
	fmt.Println(dsn, "dsn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
		os.Exit(2)
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&models.User{})

	PostgresDB = DB_Instance{
		DB: db,
	}

}
