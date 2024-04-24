package session

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateUser(context.Context, *Session) (*Session, error)
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

func (s *service) CreateUser(ctx context.Context, session *Session) (*Session, error) {
	insertedSession, err := s.repository.Insert(ctx, session)
	if err != nil {
		return nil, err
	}

	return insertedSession, nil
}
