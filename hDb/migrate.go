package helpers

import (
	"log"
)

func Migrate(models ...interface{}) error {
	err := Get().AutoMigrate(models...)
	if err != nil {
		return err
	}

	log.Println("Database migration completed successfully.")
	return nil
}
