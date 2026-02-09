package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// 1. Get ENV Variables
	err := godotenv.Load("../.env")
	if err != nil {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️  Warning: Could not load .env file, relying on System Env")
		}
	}

	// 2. Read Config
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// 3. Make Connection String
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		host, user, pass, dbname, port,
	)

	// 4. Connect
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("🔥 Failed to connect to database: ", err)
	}

	// 5. Set Global DB Variable
	DB = database
	log.Println("🚀 Database Connected Successfully!")
}
