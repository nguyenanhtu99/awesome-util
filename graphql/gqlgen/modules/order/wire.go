//go:build wireinject
// +build wireinject

package order

import (
	"awesome-util/graphql/gqlgen/modules/order/orderrepo"
	"awesome-util/graphql/gqlgen/modules/order/orderservice"

	"github.com/google/wire"
)

var (
	_ orderservice.Repository = (*orderrepo.Repository)(nil)
)

func New() (*orderservice.Service, error) {
	wire.Build(
		orderservice.New,

		wire.Bind(new(orderservice.Repository), new(*orderrepo.Repository)),
		orderrepo.New,
	)

	return &orderservice.Service{}, nil
}
