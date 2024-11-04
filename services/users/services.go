package users

import "logbun/models"

type Service interface {
	CreateUser(user models.User) (*models.User, error)
	UpdateUser(user models.User) error
	FetchProfileByEmail(email string) (*models.User, error)
}

type userSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &userSvc{repo: r}
}

// FetchProfileByEmail implements Service.
func (s *userSvc) FetchProfileByEmail(email string) (*models.User, error) {
	return s.repo.FetchProfileByEmail(email)
}

// CreateUser implements Service.
func (s *userSvc) CreateUser(user models.User) (*models.User, error) {
	return s.repo.CreateUser(user)
}

// UpdateUser implements Service.
func (s *userSvc) UpdateUser(user models.User) error {
	return s.repo.UpdateUser(user)
}
