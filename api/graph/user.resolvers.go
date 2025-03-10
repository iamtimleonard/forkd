package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"forkd/db"
	"forkd/graph/model"

	"github.com/jackc/pgx/v5/pgtype"
)

// Recipes is the resolver for the recipes field.
func (r *userResolver) Recipes(ctx context.Context, obj *model.User, limit *int, nextCursor *string) (*model.PaginatedRecipes, error) {
	var params db.ListRecipesByAuthorParams
	if obj == nil {
		return nil, fmt.Errorf("missing user object")
	}
	id := pgtype.UUID{
		Bytes: obj.ID,
		Valid: true,
	}
	params.AuthorID = id
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
	result, err := r.Queries.ListRecipesByAuthor(ctx, params)
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

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
