package helpersDatabase

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB

func New() error {
	var err error

	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, name, port)

	dsn += " TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	if db != nil {
		instance = db
		log.Print("Database connected")
	}

	return nil
}

func Get() *gorm.DB {
	if instance == nil {
		log.Println("Database connection not found!")
		max_attempts := 3
		for attempt := 0; attempt < max_attempts; attempt++ {
			log.Println("retrying connect... attempt: ", attempt)
			New()
			if instance != nil {
				return instance
			}
		}
	}
	return instance
}
