package service

import "github.com/OlegBogdanov06102017/auth-app/pkg/repository"

type Authorization interface {
}

type List interface {
}

type Service struct {
	Authorization
	List
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
