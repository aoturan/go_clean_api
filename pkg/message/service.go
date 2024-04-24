package message

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateUser(context.Context, *Message) (*Message, error)
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

func (s *service) CreateUser(ctx context.Context, message *Message) (*Message, error) {
	insertedMessage, err := s.repository.Insert(ctx, message)
	if err != nil {
		return nil, err
	}

	return insertedMessage, nil
}
