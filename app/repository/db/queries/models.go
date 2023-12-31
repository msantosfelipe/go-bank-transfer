// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package queries

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID
	Name      string
	Cpf       string
	Balance   int64
	CreatedAt time.Time
}

type Login struct {
	Cpf    string
	Secret string
}

type Transfer struct {
	ID                   uuid.UUID
	AccountOriginID      uuid.UUID
	AccountDestinationID uuid.UUID
	Amount               int64
	CreatedAt            time.Time
}
