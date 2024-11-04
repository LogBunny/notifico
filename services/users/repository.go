package users

import "logbun/models"

type Repository interface {
	CreateUser(user models.User) (*models.User, error)
	UpdateUser(user models.User) error
	FetchProfileByEmail(email string) (*models.User, error)
}
