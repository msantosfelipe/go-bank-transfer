-- name: CreateLogin :one
INSERT INTO logins(cpf, secret)
VALUES ($1, $2)
RETURNING cpf;

-- name: GetLogin :one
SELECT *
FROM logins l
WHERE l.cpf = $1;
