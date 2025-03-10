// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: revision.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getFeaturedRevisionByRecipeId = `-- name: GetFeaturedRevisionByRecipeId :one
SELECT
  recipe_revisions.id,
  recipe_revisions.recipe_id,
  recipe_revisions.parent_id,
  recipe_revisions.recipe_description,
  recipe_revisions.change_comment,
  recipe_revisions.title,
  recipe_revisions.publish_date
FROM
  recipes
JOIN recipe_revisions ON recipes.featured_revision = recipe_revisions.id
WHERE
  recipes.id = $1
LIMIT 1
`

func (q *Queries) GetFeaturedRevisionByRecipeId(ctx context.Context, id pgtype.UUID) (RecipeRevision, error) {
	row := q.db.QueryRow(ctx, getFeaturedRevisionByRecipeId, id)
	var i RecipeRevision
	err := row.Scan(
		&i.ID,
		&i.RecipeID,
		&i.ParentID,
		&i.RecipeDescription,
		&i.ChangeComment,
		&i.Title,
		&i.PublishDate,
	)
	return i, err
}

const getForkedFromRevisionByRecipeId = `-- name: GetForkedFromRevisionByRecipeId :one

SELECT
  recipe_revisions.id,
  recipe_revisions.recipe_id,
  recipe_revisions.parent_id,
  recipe_revisions.recipe_description,
  recipe_revisions.change_comment,
  recipe_revisions.title,
  recipe_revisions.publish_date
FROM
  recipes
JOIN recipe_revisions ON recipes.forked_from = recipe_revisions.id
WHERE
  recipes.id = $1
LIMIT 1
`

// Limit for pagination
func (q *Queries) GetForkedFromRevisionByRecipeId(ctx context.Context, id pgtype.UUID) (RecipeRevision, error) {
	row := q.db.QueryRow(ctx, getForkedFromRevisionByRecipeId, id)
	var i RecipeRevision
	err := row.Scan(
		&i.ID,
		&i.RecipeID,
		&i.ParentID,
		&i.RecipeDescription,
		&i.ChangeComment,
		&i.Title,
		&i.PublishDate,
	)
	return i, err
}

const getRecipeRevisionById = `-- name: GetRecipeRevisionById :one
SELECT
  id,
  recipe_id,
  parent_id,
  recipe_description,
  change_comment,
  title,
  publish_date
FROM
  recipe_revisions
WHERE
  id = $1
LIMIT 1
`

func (q *Queries) GetRecipeRevisionById(ctx context.Context, id pgtype.UUID) (RecipeRevision, error) {
	row := q.db.QueryRow(ctx, getRecipeRevisionById, id)
	var i RecipeRevision
	err := row.Scan(
		&i.ID,
		&i.RecipeID,
		&i.ParentID,
		&i.RecipeDescription,
		&i.ChangeComment,
		&i.Title,
		&i.PublishDate,
	)
	return i, err
}

const getRecipeRevisionByIngredientId = `-- name: GetRecipeRevisionByIngredientId :one
SELECT
  recipe_revisions.id,
  recipe_revisions.recipe_id,
  recipe_revisions.parent_id,
  recipe_revisions.recipe_description,
  recipe_revisions.change_comment,
  recipe_revisions.title,
  recipe_revisions.publish_date
FROM
  recipe_ingredients
JOIN
  recipe_revisions ON recipe_ingredients.revision_id = recipe_revisions.id
WHERE
  recipe_ingredients.id = $1
LIMIT 1
`

func (q *Queries) GetRecipeRevisionByIngredientId(ctx context.Context, id int64) (RecipeRevision, error) {
	row := q.db.QueryRow(ctx, getRecipeRevisionByIngredientId, id)
	var i RecipeRevision
	err := row.Scan(
		&i.ID,
		&i.RecipeID,
		&i.ParentID,
		&i.RecipeDescription,
		&i.ChangeComment,
		&i.Title,
		&i.PublishDate,
	)
	return i, err
}

const getRecipeRevisionByStepId = `-- name: GetRecipeRevisionByStepId :one
SELECT
  recipe_revisions.id,
  recipe_revisions.recipe_id,
  recipe_revisions.parent_id,
  recipe_revisions.recipe_description,
  recipe_revisions.change_comment,
  recipe_revisions.title,
  recipe_revisions.publish_date
FROM
  recipe_steps
JOIN
  recipe_revisions ON recipe_steps.revision_id = recipe_revisions.id
WHERE
  recipe_steps.id = $1
LIMIT 1
`

func (q *Queries) GetRecipeRevisionByStepId(ctx context.Context, id int64) (RecipeRevision, error) {
	row := q.db.QueryRow(ctx, getRecipeRevisionByStepId, id)
	var i RecipeRevision
	err := row.Scan(
		&i.ID,
		&i.RecipeID,
		&i.ParentID,
		&i.RecipeDescription,
		&i.ChangeComment,
		&i.Title,
		&i.PublishDate,
	)
	return i, err
}

const listRecipeRevisions = `-- name: ListRecipeRevisions :many
SELECT
  id,
  recipe_id,
  parent_id,
  recipe_description,
  change_comment,
  title,
  publish_date
FROM
  recipe_revisions
WHERE
  recipe_id = $1
  AND id > $2 -- Cursor for pagination
ORDER BY id
LIMIT $3
`

type ListRecipeRevisionsParams struct {
	RecipeID pgtype.UUID
	ID       pgtype.UUID
	Limit    int32
}

func (q *Queries) ListRecipeRevisions(ctx context.Context, arg ListRecipeRevisionsParams) ([]RecipeRevision, error) {
	rows, err := q.db.Query(ctx, listRecipeRevisions, arg.RecipeID, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RecipeRevision
	for rows.Next() {
		var i RecipeRevision
		if err := rows.Scan(
			&i.ID,
			&i.RecipeID,
			&i.ParentID,
			&i.RecipeDescription,
			&i.ChangeComment,
			&i.Title,
			&i.PublishDate,
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
