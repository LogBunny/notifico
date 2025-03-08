package emails

import "logbun/pkg/models"

type Repository interface {
	GetEmails(email string) ([]models.Email, error)
	DeleteEmails() error
}
