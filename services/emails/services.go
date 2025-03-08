package emails

import "logbun/pkg/models"

type Service interface {
	GetEmails(email string) ([]models.Email, error)
	DeleteEmails() error
}

type userSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &userSvc{repo: r}
}

// Get emails implements Service.
func (s *userSvc) GetEmails(email string) ([]models.Email, error) {
	return s.repo.GetEmails(email)
}

func (s *userSvc) DeleteEmails() error {
	return s.repo.DeleteEmails()
}
