package authapp

import "github.com/google/uuid"

type UserList struct {
	ID        int
	UserID    uuid.UUID
	CountryId int
}
