package repository

import (
	authapp "github.com/OlegBogdanov06102017/auth-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateCustomer(customer authapp.Customer) (int, error)
}

type List interface {
}

type Repository struct {
	Authorization
	List
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
