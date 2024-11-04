package db

import (
	"logbun/services/users"
)

var (
	UsersSvc users.Service = nil
)

func InitServices() {
	db := GetDB()
	usersRepo := users.NewPostgresRepo(db)
	UsersSvc = users.NewService(usersRepo)

}
