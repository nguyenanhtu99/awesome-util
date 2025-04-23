package userservice

import (
	"awesome-util/graphql/go-graphql/modules/user/usermodel"
	"awesome-util/graphql/go-graphql/modules/user/userrepo"
	"context"
	"fmt"
)

type Service struct {
	repo *userrepo.Repository
}

func New() *Service {
	return &Service{repo: userrepo.New()}
}

func (s *Service) Create(ctx context.Context, name string) (*usermodel.User, error) {
	user := usermodel.User{Name: name}
	if err := s.repo.Create(ctx, &user); err != nil {
		return nil, fmt.Errorf("failed to create user, err: %w", err)
	}

	return &user, nil
}

func (s *Service) Get(ctx context.Context, id string) (*usermodel.User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user, err: %w", err)
	}

	return user, nil
}
