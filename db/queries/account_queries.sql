-- name: CountAccountByCpf :one
SELECT count(1) 
FROM accounts a 
WHERE a.cpf = $1;

-- name: CreateAccount :one
INSERT INTO accounts(id, name, cpf)
VALUES ($1, $2, $3)
RETURNING id;