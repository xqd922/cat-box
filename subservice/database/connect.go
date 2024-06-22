package database

import (
	"github.com/daifiyum/cat-box/config"
	"github.com/daifiyum/cat-box/subservice/models"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() error {
	var err error

	DB, err = gorm.Open(sqlite.Open(config.Config("DB_PATH")), &gorm.Config{})

	if err != nil {
		log.Error("failed to connect database")
		return err
	}

	// migrate database
	DB.AutoMigrate(&models.Subscriptions{}, &models.Users{}, &models.Options{})

	// create default options if not exists
	options := models.Options{}
	DB.FirstOrCreate(&options)
	return nil
}
