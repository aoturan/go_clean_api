package user

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateUser(context.Context, *User) (*User, error)
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

func (s *service) CreateUser(ctx context.Context, user *User) (*User, error) {
	insertedUser, err := s.repository.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	return insertedUser, nil
}
