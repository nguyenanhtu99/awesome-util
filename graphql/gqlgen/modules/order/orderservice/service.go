package orderservice

import (
	"awesome-util/graphql/gqlgen/modules/order/ordermodel"
	"context"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, order *ordermodel.Order) error
	Get(ctx context.Context, id string) (*ordermodel.Order, error)
	GetByUserId(ctx context.Context, userId string) ([]*ordermodel.Order, error)
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

func (s *Service) Create(ctx context.Context, userId string) (*ordermodel.Order, error) {
	order := ordermodel.Order{UserID: userId}
	if err := s.repo.Create(ctx, &order); err != nil {
		return nil, fmt.Errorf("failed to create order, err: %w", err)
	}

	return &order, nil
}

func (s *Service) Get(ctx context.Context, id string) (*ordermodel.Order, error) {
	fmt.Println("Getting order")
	order, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order, err: %w", err)
	}

	return order, nil
}

func (s *Service) GetOwn(ctx context.Context, userId string) ([]*ordermodel.Order, error) {
	fmt.Println("Getting own orders")
	orders, err := s.repo.GetByUserId(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get own orders, err: %w", err)
	}

	return orders, nil
}
