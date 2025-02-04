package repository

import (
	"fmt"

	authapp "github.com/OlegBogdanov06102017/auth-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateCustomer(customer authapp.Customer) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (phone, email) values ($1, $2) RETURNING id", customersTable)

	row := r.db.QueryRow(query, customer.Phone, customer.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
