-- name: CreateTransfer :one
INSERT INTO transfers(id, account_origin_id, account_destination_id, amount)
VALUES ($1, $2, $3, $4)
RETURNING id;
