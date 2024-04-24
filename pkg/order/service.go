package order

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateUser(context.Context, *Order) (*Order, error)
}

type service struct {
	repository Repository
}

func NewService(client *mongo.Client, dbName string) Service {
	repository := NewRepository(client, dbName)
	return &service{
		repository: repository,
	}
}

func (s *service) CreateUser(ctx context.Context, order *Order) (*Order, error) {
	insertedOrder, err := s.repository.Insert(ctx, order)
	if err != nil {
		return nil, err
	}

	return insertedOrder, nil
}
