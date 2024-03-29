package service

import (
	"alvile-api/models"
	"alvile-api/pkg/repository"
)

type Registration interface {
	CreateUser(user models.AccountInput) (*models.AccountRegistrationOutput, error)
	CreateSession(session *models.Session, email, password string) (*models.SessionOutput, error)
	CreateScheme(scheme models.Scheme, email string) (*models.SchemeOutput, error)
	CheckAuthorization(hed string) (string, error)
	GetScheme(email string) ([]models.SchemeOutput, error)
}

type Service struct {
	Registration
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Registration: NewAuthService(repos.Registration),
	}
}
