-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserWithUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserWithEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserWithEmailAndPassword :one
SELECT * FROM users
WHERE email = $1 AND password = $2 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
    username, email, password, balance
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET 
    username = $2,
    email = $3,
    password = $4,
    balance = $5
WHERE id = $1
RETURNING *;

-- name: UpdateUserBalance :one
UPDATE users
SET 
    balance = $2
WHERE id = $1
RETURNING *;