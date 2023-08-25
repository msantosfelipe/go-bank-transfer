-- name: CountAccountByCpf :one
SELECT count(1) 
FROM accounts a 
WHERE a.cpf = $1;

-- name: CreateAccount :one
INSERT INTO accounts(id, name, cpf)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetAccounts :many
SELECT a.*, l.secret
FROM accounts a
INNER JOIN logins l ON a.cpf = l.cpf;

-- name: GetAccountsBalances :many
SELECT a.id, a.balance
FROM accounts a
WHERE a.id = ANY($1::uuid[]);

-- name: UpdateAccountBalance :exec
UPDATE accounts a
SET balance = $2
WHERE a.id = $1;
