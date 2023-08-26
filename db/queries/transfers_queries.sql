-- name: CreateTransfer :one
INSERT INTO transfers(id, account_origin_id, account_destination_id, amount)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetTransfers :many
SELECT *
FROM transfers t
WHERE t.account_origin_id = $1;
