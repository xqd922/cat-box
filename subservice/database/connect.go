package database

import (
	"github.com/daifiyum/cat-box/config"
	"github.com/daifiyum/cat-box/subservice/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.Config("DB_PATH")), &gorm.Config{})

	if err != nil {
		return err
	}

	// migrate database
	DB.AutoMigrate(&models.Subscriptions{}, &models.Users{}, &models.Options{})

	// create default options if not exists
	options := models.Options{}
	DB.FirstOrCreate(&options)
	return nil
}
