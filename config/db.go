package config

import (
	"fmt"
	"log"
	"os"

	"github.com/bintangyosua/shortenurl/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set!")
	}
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
	}
	DB.AutoMigrate(&models.URL{})
	fmt.Println("Connected to DB!")
}

