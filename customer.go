package authapp

import "github.com/google/uuid"

type Customer struct {
	ID    uuid.UUID `json:"id"`
	Phone string    `json:"phone"`
	Email string    `json:"email"`
}
