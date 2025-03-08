package migrations

import (
	"logbun/db"
	"logbun/pkg/models"
)

func Migrate() {
	database := db.GetDB()
	database.AutoMigrate(&models.Email{})
}
