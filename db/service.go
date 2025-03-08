package db

import (
	"logbun/services/emails"
	"logbun/services/users"
)

var (
	UsersSvc users.Service  = nil
	EmailSvc emails.Service = nil
)

func InitServices() {
	db := GetDB()
	usersRepo := users.NewPostgresRepo(db)
	UsersSvc = users.NewService(usersRepo)

	emailRepo := emails.NewPostgresRepo(db)
	EmailSvc = emails.NewService(emailRepo)
}
