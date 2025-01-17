// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
	display_name,
	email
) VALUES (
	$1,
	$2
)
RETURNING users.id, users.display_name, users.email, users.join_date, users.updated_at
`

type CreateUserParams struct {
	DisplayName string
	Email       string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.DisplayName, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.DisplayName,
		&i.Email,
		&i.JoinDate,
		&i.UpdatedAt,
	)
	return i, err
}

const getAuthorByRecipeId = `-- name: GetAuthorByRecipeId :one
SELECT
  users.id,
  users.display_name,
  users.email,
  users.join_date,
  users.updated_at
FROM
  users
JOIN recipes ON users.id = recipes.author_id
WHERE
  recipes.id = $1
LIMIT 1
`

func (q *Queries) GetAuthorByRecipeId(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getAuthorByRecipeId, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.DisplayName,
		&i.Email,
		&i.JoinDate,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT users.id, users.display_name, users.email, users.join_date, users.updated_at FROM users WHERE users.email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.DisplayName,
		&i.Email,
		&i.JoinDate,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT users.id, users.display_name, users.email, users.join_date, users.updated_at FROM users WHERE users.id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.DisplayName,
		&i.Email,
		&i.JoinDate,
		&i.UpdatedAt,
	)
	return i, err
}
