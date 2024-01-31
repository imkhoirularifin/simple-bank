// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username, email, password, balance
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, username, email, password, balance
`

type CreateUserParams struct {
	Username string
	Email    string
	Password string
	Balance  int64
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Balance,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Balance,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, password, balance FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Balance,
	)
	return i, err
}

const getUserWithEmail = `-- name: GetUserWithEmail :one
SELECT id, username, email, password, balance FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserWithEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserWithEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Balance,
	)
	return i, err
}

const getUserWithEmailAndPassword = `-- name: GetUserWithEmailAndPassword :one
SELECT id, username, email, password, balance FROM users
WHERE email = $1 AND password = $2 LIMIT 1
`

type GetUserWithEmailAndPasswordParams struct {
	Email    string
	Password string
}

func (q *Queries) GetUserWithEmailAndPassword(ctx context.Context, arg GetUserWithEmailAndPasswordParams) (User, error) {
	row := q.db.QueryRow(ctx, getUserWithEmailAndPassword, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Balance,
	)
	return i, err
}

const getUserWithUsername = `-- name: GetUserWithUsername :one
SELECT id, username, email, password, balance FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserWithUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserWithUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Balance,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, password, balance FROM users
ORDER BY username
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.Balance,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET 
    username = $2,
    email = $3,
    password = $4,
    balance = $5
WHERE id = $1
RETURNING id, username, email, password, balance
`

type UpdateUserParams struct {
	ID       pgtype.UUID
	Username string
	Email    string
	Password string
	Balance  int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Balance,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Balance,
	)
	return i, err
}

const updateUserBalance = `-- name: UpdateUserBalance :one
UPDATE users
SET 
    balance = $2
WHERE id = $1
RETURNING id, username, email, password, balance
`

type UpdateUserBalanceParams struct {
	ID      pgtype.UUID
	Balance int64
}

func (q *Queries) UpdateUserBalance(ctx context.Context, arg UpdateUserBalanceParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserBalance, arg.ID, arg.Balance)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Balance,
	)
	return i, err
}