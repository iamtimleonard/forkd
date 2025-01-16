// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: recipe.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRecipe = `-- name: CreateRecipe :one
INSERT INTO recipes (
  author_id,
  forked_from,
  slug,
  private
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
`

type CreateRecipeParams struct {
	AuthorID   int64
	ForkedFrom pgtype.Int8
	Slug       string
	Private    bool
}

func (q *Queries) CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error) {
	row := q.db.QueryRow(ctx, createRecipe,
		arg.AuthorID,
		arg.ForkedFrom,
		arg.Slug,
		arg.Private,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Slug,
		&i.Private,
		&i.InitialPublishDate,
		&i.ForkedFrom,
		&i.FeaturedRevision,
	)
	return i, err
}

const getRecipeById = `-- name: GetRecipeById :one
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
FROM
  recipes
WHERE
  id = $1
LIMIT 1
`

func (q *Queries) GetRecipeById(ctx context.Context, id int64) (Recipe, error) {
	row := q.db.QueryRow(ctx, getRecipeById, id)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Slug,
		&i.Private,
		&i.InitialPublishDate,
		&i.ForkedFrom,
		&i.FeaturedRevision,
	)
	return i, err
}

const getRecipeByRevisionID = `-- name: GetRecipeByRevisionID :one
SELECT
  recipes.id,
  recipes.author_id,
  recipes.slug,
  recipes.private,
  recipes.initial_publish_date,
  recipes.forked_from,
  recipes.featured_revision
FROM
  recipe_revisions
JOIN
  recipes ON recipe_revisions.recipe_id = recipes.id
WHERE
  recipes.id = $1
LIMIT 1
`

func (q *Queries) GetRecipeByRevisionID(ctx context.Context, id int64) (Recipe, error) {
	row := q.db.QueryRow(ctx, getRecipeByRevisionID, id)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Slug,
		&i.Private,
		&i.InitialPublishDate,
		&i.ForkedFrom,
		&i.FeaturedRevision,
	)
	return i, err
}

const getRecipeBySlug = `-- name: GetRecipeBySlug :one
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
FROM
  recipes
WHERE
  slug = $1
LIMIT 1
`

func (q *Queries) GetRecipeBySlug(ctx context.Context, slug string) (Recipe, error) {
	row := q.db.QueryRow(ctx, getRecipeBySlug, slug)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Slug,
		&i.Private,
		&i.InitialPublishDate,
		&i.ForkedFrom,
		&i.FeaturedRevision,
	)
	return i, err
}

const getRecipeRevisionByParentID = `-- name: GetRecipeRevisionByParentID :one
SELECT
  parent.id,
  parent.recipe_id,
  parent.parent_id,
  parent.recipe_description,
  parent.change_comment,
  parent.title,
  parent.publish_date
FROM
  recipe_revisions child
JOIN
  recipe_revisions parent ON child.parent_id = parent.id
WHERE
  recipe_revisions.id = $1
LIMIT 1
`

func (q *Queries) GetRecipeRevisionByParentID(ctx context.Context, id int64) (RecipeRevision, error) {
	row := q.db.QueryRow(ctx, getRecipeRevisionByParentID, id)
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

const listIngredientsByRecipeRevisionID = `-- name: ListIngredientsByRecipeRevisionID :many
SELECT
  recipe_ingredients.id,
  recipe_ingredients.revision_id,
  recipe_ingredients.ingredient,
  recipe_ingredients.quantity,
  recipe_ingredients.unit,
  recipe_ingredients.comment
FROM
  recipe_revisions
JOIN
  recipe_ingredients ON recipe_revisions.id = recipe_ingredients.revision_id 
WHERE
  recipe_revisions.id = $1
ORDER BY recipe_ingredients.id
`

func (q *Queries) ListIngredientsByRecipeRevisionID(ctx context.Context, id int64) ([]RecipeIngredient, error) {
	rows, err := q.db.Query(ctx, listIngredientsByRecipeRevisionID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RecipeIngredient
	for rows.Next() {
		var i RecipeIngredient
		if err := rows.Scan(
			&i.ID,
			&i.RevisionID,
			&i.Ingredient,
			&i.Quantity,
			&i.Unit,
			&i.Comment,
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

const listRecipes = `-- name: ListRecipes :many
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
FROM
  recipes
WHERE
  id > $1
ORDER BY id
LIMIT $2
`

type ListRecipesParams struct {
	ID    int64
	Limit int32
}

func (q *Queries) ListRecipes(ctx context.Context, arg ListRecipesParams) ([]Recipe, error) {
	rows, err := q.db.Query(ctx, listRecipes, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Slug,
			&i.Private,
			&i.InitialPublishDate,
			&i.ForkedFrom,
			&i.FeaturedRevision,
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

const listRecipesByAuthor = `-- name: ListRecipesByAuthor :many
SELECT
  id,
  author_id,
  slug,
  private,
  initial_publish_date,
  forked_from,
  featured_revision
FROM
  recipes
WHERE
  author_id = $1 AND id > $2
ORDER BY id
LIMIT $3
`

type ListRecipesByAuthorParams struct {
	AuthorID int64
	ID       int64
	Limit    int32
}

func (q *Queries) ListRecipesByAuthor(ctx context.Context, arg ListRecipesByAuthorParams) ([]Recipe, error) {
	rows, err := q.db.Query(ctx, listRecipesByAuthor, arg.AuthorID, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Slug,
			&i.Private,
			&i.InitialPublishDate,
			&i.ForkedFrom,
			&i.FeaturedRevision,
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

const listStepsByRecipeRevisionID = `-- name: ListStepsByRecipeRevisionID :many
SELECT
  recipe_steps.id,
  recipe_steps.revision_id,
  recipe_steps.content,
  recipe_steps.index
FROM
  recipe_revisions
JOIN
  recipe_steps ON recipe_revisions.id = recipe_steps.revision_id
WHERE
  recipe_revisions.id = $1
ORDER BY
  recipe_steps.id
`

func (q *Queries) ListStepsByRecipeRevisionID(ctx context.Context, id int64) ([]RecipeStep, error) {
	rows, err := q.db.Query(ctx, listStepsByRecipeRevisionID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RecipeStep
	for rows.Next() {
		var i RecipeStep
		if err := rows.Scan(
			&i.ID,
			&i.RevisionID,
			&i.Content,
			&i.Index,
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
