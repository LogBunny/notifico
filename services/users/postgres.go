package users

import (
	"logbun/pkg/models"

	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

// CreateUser implements Repository.
func (r *repo) CreateUser(user models.User) (*models.User, error) {

	err := r.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FetchProfileByEmail implements Repository.
func (r *repo) FetchProfileByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.DB.First(user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser implements Repository.
func (r *repo) UpdateUser(user models.User) error {
	return r.DB.Updates(user).Error
}
