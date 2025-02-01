package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type List interface {
}

type Repository struct {
	Authorization
	List
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
