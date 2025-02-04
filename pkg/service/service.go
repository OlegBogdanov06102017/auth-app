package service

import (
	authapp "github.com/OlegBogdanov06102017/auth-app"
	"github.com/OlegBogdanov06102017/auth-app/pkg/repository"
)

type Authorization interface {
	CreateCustomer(customer authapp.Customer) (int, error)
}

type List interface {
}

type Service struct {
	Authorization
	List
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
