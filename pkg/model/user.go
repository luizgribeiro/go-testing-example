package model

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Name  string    `db:"name" json:"name"`
	Email string    `db:"email" json:"email"`
}

type UserAccount struct {
	User
	Account
	AccountID uuid.UUID `json:"accountId"`
}
