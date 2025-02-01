package authapp

import "github.com/google/uuid"

type CustomerList struct {
	ID        int
	UserID    uuid.UUID
	CountryId int
}
