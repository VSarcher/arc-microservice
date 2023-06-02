package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/VSarcher/arc-microservice/auth-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB_Instance struct {
	DB *gorm.DB
}

var PostgresDB DB_Instance

func ConnectDB() error {
	dsn := fmt.Sprintf("host=postgres-db user=%s password=%s dbname=%s sslmode=disable TimeZone=EST",
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
		return errors.New("Failed to connect to database")
	}
	fmt.Println("Pass")
	db.Logger = logger.Default.LogMode(logger.Info)

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
		return errors.New("Error on migrating db schema!")
	}
	// fmt.Println("Passed!", err)
	PostgresDB = DB_Instance{
		DB: db,
	}
	return nil
}
