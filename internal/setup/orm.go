package setup

import (
	"github.com/scriptibus/moinvoice/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitORM() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.Customer{},
		&models.Project{},
		&models.Booking{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
