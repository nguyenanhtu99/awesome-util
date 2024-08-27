//go:build wireinject
// +build wireinject

package user

import (
	"awesome-util/graphql/gqlgen/modules/user/userrepo"
	"awesome-util/graphql/gqlgen/modules/user/userservice"

	"github.com/google/wire"
)

var (
	_ userservice.Repository = (*userrepo.Repository)(nil)
)

func New() (*userservice.Service, error) {
	wire.Build(
		userservice.New,

		wire.Bind(new(userservice.Repository), new(*userrepo.Repository)),
		userrepo.New,
	)

	return &userservice.Service{}, nil
}
