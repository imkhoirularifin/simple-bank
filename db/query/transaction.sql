-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: ListTransactions :many
SELECT * FROM transactions
ORDER BY date DESC;

-- name: CreateTransaction :one
INSERT INTO transactions (
    from_user_id, to_user_id, amount, description
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;