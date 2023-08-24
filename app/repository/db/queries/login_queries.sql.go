// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: login_queries.sql

package queries

import (
	"context"

	"github.com/google/uuid"
)

const createLogin = `-- name: CreateLogin :one
INSERT INTO logins(cpf, secret)
VALUES ($1, $2)
RETURNING cpf
`

type CreateLoginParams struct {
	Cpf    string
	Secret string
}

func (q *Queries) CreateLogin(ctx context.Context, arg CreateLoginParams) (string, error) {
	row := q.db.QueryRow(ctx, createLogin, arg.Cpf, arg.Secret)
	var cpf string
	err := row.Scan(&cpf)
	return cpf, err
}

const getLoginAndAccount = `-- name: GetLoginAndAccount :one
SELECT l.cpf, l.secret, a.id 
FROM logins l
INNER JOIN accounts a ON a.cpf = l.cpf 
WHERE l.cpf = $1
`

type GetLoginAndAccountRow struct {
	Cpf    string
	Secret string
	ID     uuid.UUID
}

func (q *Queries) GetLoginAndAccount(ctx context.Context, cpf string) (GetLoginAndAccountRow, error) {
	row := q.db.QueryRow(ctx, getLoginAndAccount, cpf)
	var i GetLoginAndAccountRow
	err := row.Scan(&i.Cpf, &i.Secret, &i.ID)
	return i, err
}
