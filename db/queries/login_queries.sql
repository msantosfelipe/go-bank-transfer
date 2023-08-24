-- name: CreateLogin :one
INSERT INTO logins(cpf, secret)
VALUES ($1, $2)
RETURNING cpf;

-- name: GetLoginAndAccount :one
SELECT l.*, a.id 
FROM logins l
INNER JOIN accounts a ON a.cpf = l.cpf 
WHERE l.cpf = $1;
