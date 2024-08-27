package graph

//go:generate go get github.com/99designs/gqlgen
//go:generate go run github.com/99designs/gqlgen generate

import (
	"awesome-util/graphql/gqlgen/modules/order/ordermodel"
	"awesome-util/graphql/gqlgen/modules/user/usermodel"
	"context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type UserService interface {
	Create(ctx context.Context, name string) (*usermodel.User, error)
	Get(ctx context.Context, id string) (*usermodel.User, error)
}

type OrderService interface {
	Create(ctx context.Context, name string) (*ordermodel.Order, error)
	Get(ctx context.Context, id string) (*ordermodel.Order, error)
	GetOwn(ctx context.Context, userId string) ([]*ordermodel.Order, error)
}

func NewResolver(userSvc UserService, orderSvc OrderService) *Resolver {
	return &Resolver{
		userSvc:  userSvc,
		orderSvc: orderSvc,
	}
}

type Resolver struct {
	userSvc  UserService
	orderSvc OrderService
}
