package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"awesome-util/graphql/gqlgen/modules/user/usermodel"
	"context"
)

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *usermodel.User) (string, error) {
	return obj.ID.Hex(), nil
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }