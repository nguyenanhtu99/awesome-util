// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package user

import (
	"awesome-util/graphql/gqlgen/modules/user/userrepo"
	"awesome-util/graphql/gqlgen/modules/user/userservice"
)

// Injectors from wire.go:

func New() (*userservice.Service, error) {
	repository := userrepo.New()
	service := userservice.New(repository)
	return service, nil
}

// wire.go:

var (
	_ userservice.Repository = (*userrepo.Repository)(nil)
)
