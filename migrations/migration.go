package migrations

import (
	"logbun/db"
	"logbun/models"
)

func Migrate() {
	database := db.GetDB()
	database.AutoMigrate(&models.User{})
}
