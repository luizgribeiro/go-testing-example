package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Name  string    `db:"name" json:"name"`
	Email string    `db:"email" json:"email"`
	BirthDate time.Time `db:"birthdate" json:"birthdate"`
}

type UserAccount struct {
	User
	Account
	AccountID uuid.UUID `json:"accountId"`
}
