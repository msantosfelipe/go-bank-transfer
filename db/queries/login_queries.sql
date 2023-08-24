-- name: CreateLogin :one
INSERT INTO logins(cpf, secret)
VALUES ($1, $2)
RETURNING cpf;