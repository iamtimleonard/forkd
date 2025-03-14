package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"forkd/db"
	"forkd/graph/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.UserQuery, error) {
	return &model.UserQuery{}, nil
}

// Recipe is the resolver for the recipe field.
func (r *queryResolver) Recipe(ctx context.Context) (*model.RecipeQuery, error) {
	return &model.RecipeQuery{}, nil
}

// ByID is the resolver for the byId field.
func (r *recipeQueryResolver) ByID(ctx context.Context, obj *model.RecipeQuery, id uuid.UUID) (*model.Recipe, error) {
	pgId := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
	result, err := r.Queries.GetRecipeById(ctx, pgId)
	return handleNoRowsOnNullableType(result, err, model.RecipeFromDBType)
}

// BySlug is the resolver for the bySlug field.
func (r *recipeQueryResolver) BySlug(ctx context.Context, obj *model.RecipeQuery, slug string) (*model.Recipe, error) {
	// Fetch the recipe by the slug
	result, err := r.Queries.GetRecipeBySlug(ctx, slug)
	return handleNoRowsOnNullableType(result, err, model.RecipeFromDBType)
}

// List is the resolver for the list field.
func (r *recipeQueryResolver) List(ctx context.Context, obj *model.RecipeQuery, limit *int, nextCursor *string) (*model.PaginatedRecipes, error) {
	var params db.ListRecipesParams
	if limit != nil {
		params.Limit = int32(*limit)
	} else {
		params.Limit = 20
	}
	if nextCursor != nil {
		cursor := new(ListRecipesCursor)
		err := cursor.Decode(*nextCursor)
		if err != nil {
			return nil, err
		}
		if !cursor.Validate(*limit) {
			return nil, fmt.Errorf("limit param does not match cursor. Limit: %d, Cursor: %d", params.Limit, cursor.Limit)
		}
		params.ID = cursor.Id
	}
	result, err := r.Queries.ListRecipes(ctx, params)
	// If there was an error, early return with the error
	if err != nil {
		return nil, err
	}
	count := len(result)
	recipes := make([]*model.Recipe, count)
	for i, recipe := range result {
		recipes[i] = model.RecipeFromDBType(recipe)
	}

	var NextCursor *string = nil

	if count == int(params.Limit) {
		cursor := ListRecipesCursor{
			Id:    result[count-1].ID,
			Limit: int(params.Limit),
		}
		encoded, err := cursor.Encode()

		if err != nil {
			return nil, err
		}

		NextCursor = &encoded
	}

	paginationInfo := model.PaginationInfo{
		Count:      count,
		NextCursor: NextCursor,
	}

	paginated := model.PaginatedRecipes{
		Items:      recipes,
		Pagination: &paginationInfo,
	}

	return &paginated, nil
}

// ByID is the resolver for the byId field.
func (r *userQueryResolver) ByID(ctx context.Context, obj *model.UserQuery, id uuid.UUID) (*model.User, error) {
	pgId := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
	result, err := r.Queries.GetUserById(ctx, pgId)
	return handleNoRowsOnNullableType(result, err, model.UserFromDBType)
}

// ByEmail is the resolver for the byEmail field.
func (r *userQueryResolver) ByEmail(ctx context.Context, obj *model.UserQuery, email string) (*model.User, error) {
	result, err := r.Queries.GetUserByEmail(ctx, email)
	return handleNoRowsOnNullableType(result, err, model.UserFromDBType)
}

// Current is the resolver for the current field.
func (r *userQueryResolver) Current(ctx context.Context, obj *model.UserQuery) (*model.User, error) {
	user, _ := r.Auth.GetUserSessionFromCtx(ctx)
	return model.UserFromDBType(*user), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// RecipeQuery returns RecipeQueryResolver implementation.
func (r *Resolver) RecipeQuery() RecipeQueryResolver { return &recipeQueryResolver{r} }

// UserQuery returns UserQueryResolver implementation.
func (r *Resolver) UserQuery() UserQueryResolver { return &userQueryResolver{r} }

type queryResolver struct{ *Resolver }
type recipeQueryResolver struct{ *Resolver }
type userQueryResolver struct{ *Resolver }
