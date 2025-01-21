// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMagicLink = `-- name: CreateMagicLink :one
INSERT INTO
  magic_links (
    user_id,
    token,
    expiry
  )
VALUES (
  $1,
  $2,
  $3
)
RETURNING
  magic_links.id,
  magic_links.token
`

type CreateMagicLinkParams struct {
	UserID pgtype.UUID
	Token  pgtype.UUID
	Expiry pgtype.Timestamp
}

type CreateMagicLinkRow struct {
	ID    pgtype.UUID
	Token pgtype.UUID
}

func (q *Queries) CreateMagicLink(ctx context.Context, arg CreateMagicLinkParams) (CreateMagicLinkRow, error) {
	row := q.db.QueryRow(ctx, createMagicLink, arg.UserID, arg.Token, arg.Expiry)
	var i CreateMagicLinkRow
	err := row.Scan(&i.ID, &i.Token)
	return i, err
}

const createSession = `-- name: CreateSession :one
WITH sesh AS (
  INSERT INTO
    sessions (
      user_id,
      expiry
    )
  VALUES (
    $1,
    $2
  )
  RETURNING
    sessions.id,
    sessions.user_id
)
SELECT users.id, users.display_name, users.email, users.join_date, users.updated_at, sesh.id FROM sesh INNER JOIN users ON sesh.user_id = users.id
`

type CreateSessionParams struct {
	UserID pgtype.UUID
	Expiry pgtype.Timestamp
}

type CreateSessionRow struct {
	User User
	ID   pgtype.UUID
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (CreateSessionRow, error) {
	row := q.db.QueryRow(ctx, createSession, arg.UserID, arg.Expiry)
	var i CreateSessionRow
	err := row.Scan(
		&i.User.ID,
		&i.User.DisplayName,
		&i.User.Email,
		&i.User.JoinDate,
		&i.User.UpdatedAt,
		&i.ID,
	)
	return i, err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM
  sessions
WHERE sessions.id = $1
`

func (q *Queries) DeleteSession(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteSession, id)
	return err
}

const getMagicLink = `-- name: GetMagicLink :one
SELECT
  magic_links.id,
  magic_links.token,
  magic_links.user_id,
  magic_links.expiry
FROM
  magic_links
WHERE
  magic_links.id = $1 AND magic_links.token = $2
LIMIT 1
`

type GetMagicLinkParams struct {
	ID    pgtype.UUID
	Token pgtype.UUID
}

type GetMagicLinkRow struct {
	ID     pgtype.UUID
	Token  pgtype.UUID
	UserID pgtype.UUID
	Expiry pgtype.Timestamp
}

func (q *Queries) GetMagicLink(ctx context.Context, arg GetMagicLinkParams) (GetMagicLinkRow, error) {
	row := q.db.QueryRow(ctx, getMagicLink, arg.ID, arg.Token)
	var i GetMagicLinkRow
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.UserID,
		&i.Expiry,
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

func (q *Queries) GetUserById(ctx context.Context, id pgtype.UUID) (User, error) {
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

const getUserBySessionId = `-- name: GetUserBySessionId :one
SELECT
  users.id, users.display_name, users.email, users.join_date, users.updated_at,
  sessions.id, sessions.user_id, sessions.expiry
FROM
  sessions
JOIN
  users ON users.id = sessions.user_id
WHERE
  sessions.id = $1
LIMIT 1
`

type GetUserBySessionIdRow struct {
	User    User
	Session Session
}

func (q *Queries) GetUserBySessionId(ctx context.Context, id pgtype.UUID) (GetUserBySessionIdRow, error) {
	row := q.db.QueryRow(ctx, getUserBySessionId, id)
	var i GetUserBySessionIdRow
	err := row.Scan(
		&i.User.ID,
		&i.User.DisplayName,
		&i.User.Email,
		&i.User.JoinDate,
		&i.User.UpdatedAt,
		&i.Session.ID,
		&i.Session.UserID,
		&i.Session.Expiry,
	)
	return i, err
}

const upsertUser = `-- name: UpsertUser :one
WITH upsert AS (
  INSERT INTO
    users (
      email,
      display_name
    )
  VALUES (
    $1,
    $2
  )
  ON CONFLICT (email)
  DO NOTHING
  RETURNING
    users.id,
    users.display_name,
    users.email,
    users.join_date,
    users.updated_at
)
SELECT
  upsert.id,
	upsert.display_name,
	upsert.email,
	upsert.join_date,
	upsert.updated_at
FROM
  upsert
UNION
SELECT
  users.id,
	users.display_name,
	users.email,
	users.join_date,
	users.updated_at
FROM
  users
WHERE
  users.email = $1
`

type UpsertUserParams struct {
	Email       string
	DisplayName string
}

type UpsertUserRow struct {
	ID          pgtype.UUID
	DisplayName string
	Email       string
	JoinDate    pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) UpsertUser(ctx context.Context, arg UpsertUserParams) (UpsertUserRow, error) {
	row := q.db.QueryRow(ctx, upsertUser, arg.Email, arg.DisplayName)
	var i UpsertUserRow
	err := row.Scan(
		&i.ID,
		&i.DisplayName,
		&i.Email,
		&i.JoinDate,
		&i.UpdatedAt,
	)
	return i, err
}
