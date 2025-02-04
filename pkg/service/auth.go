package service

import (
	authapp "github.com/OlegBogdanov06102017/auth-app"
	"github.com/OlegBogdanov06102017/auth-app/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateCustomer(customer authapp.Customer) (int, error) {
	return s.repo.CreateCustomer(customer)
}
