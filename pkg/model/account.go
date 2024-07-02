package model

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Handler   string    `db:"handler" json:"handler"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	LastLogin time.Time `db:"last_login" json:"last_login"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
}
