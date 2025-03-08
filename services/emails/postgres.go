package emails

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

// FetchProfileByEmail implements Repository.
func (r *repo) GetEmails(email string) ([]models.Email, error) {
	emails := make([]models.Email, 0)
	err := r.DB.Find(&emails, "emails.to = ?", email).Error
	if err != nil {
		return nil, err
	}
	return emails, nil
}

// UpdateUser implements Repository.
func (r *repo) DeleteEmails() error {
	return r.DB.Model(&models.Email{}).Delete(&models.Email{}).Where("created_at < NOW() - INTERVAL '12 hours'").Error
}
