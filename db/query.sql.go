// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, email, name, password) VALUES (?, ?, ?, ?)
`

type CreateUserParams struct {
	ID       string
	Email    string
	Name     string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Email,
		arg.Name,
		arg.Password,
	)
	return err
}

const getUserWithEmail = `-- name: GetUserWithEmail :one
SELECT id, email, password, name FROM users WHERE email = ? LIMIT 1
`

func (q *Queries) GetUserWithEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserWithEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
	)
	return i, err
}
