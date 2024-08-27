package userservice

import (
	"awesome-util/graphql/gqlgen/modules/user/usermodel"
	"context"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, user *usermodel.User) error
	Get(ctx context.Context, id string) (*usermodel.User, error)
}

type Service struct {
	repo Repository
}

func New(
	repo Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, name string) (*usermodel.User, error) {
	user := usermodel.User{Name: name}
	if err := s.repo.Create(ctx, &user); err != nil {
		return nil, fmt.Errorf("failed to create user, err: %v", err)
	}

	return &user, nil
}

func (s *Service) Get(ctx context.Context, id string) (*usermodel.User, error) {
	fmt.Println("Getting user")
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user, err: %v", err)
	}

	return user, nil
}
