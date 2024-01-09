package database

import (
	"com/beeda/search/app/models"
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.TestModel{},
	)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
