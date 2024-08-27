//go:build wireinject
// +build wireinject

package graph

import (
	"awesome-util/graphql/gqlgen/modules/order"
	"awesome-util/graphql/gqlgen/modules/order/orderservice"
	"awesome-util/graphql/gqlgen/modules/user"
	"awesome-util/graphql/gqlgen/modules/user/userservice"

	"github.com/google/wire"
)

var (
	_ UserService  = (*userservice.Service)(nil)
	_ OrderService = (*orderservice.Service)(nil)
)

func New() (*Resolver, error) {
	wire.Build(
		NewResolver,

		wire.Bind(new(UserService), new(*userservice.Service)),
		user.New,

		wire.Bind(new(OrderService), new(*orderservice.Service)),
		order.New,
	)

	return &Resolver{}, nil
}
