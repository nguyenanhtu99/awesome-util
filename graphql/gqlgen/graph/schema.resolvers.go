package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"awesome-util/graphql/gqlgen/modules/order/ordermodel"
	"awesome-util/graphql/gqlgen/modules/user/usermodel"
	"context"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*usermodel.User, error) {
	return r.userSvc.Create(ctx, name)
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, userID string) (*ordermodel.Order, error) {
	return r.orderSvc.Create(ctx, userID)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*usermodel.User, error) {
	return r.userSvc.Get(ctx, id)
}

// Order is the resolver for the order field.
func (r *queryResolver) Order(ctx context.Context, id string) (*ordermodel.Order, error) {
	return r.orderSvc.Get(ctx, id)
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context, userID string) ([]*ordermodel.Order, error) {
	return r.orderSvc.GetOwn(ctx, userID)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
