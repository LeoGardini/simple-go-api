package migrations

import (
	"github.com/LeoGardini/simple-go-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
