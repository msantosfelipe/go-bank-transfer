// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package queries

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type Account struct {
	ID        uuid.UUID
	Name      string
	Cpf       string
	Balance   pgtype.Numeric
	CreatedAt time.Time
}

type Login struct {
	Cpf    string
	Secret string
}