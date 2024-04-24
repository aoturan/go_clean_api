package session

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const coll = "sessions"

type Repository interface {
	Insert(context.Context, *Session) (*Session, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client, dbName string) Repository {
	return &repository{
		collection: client.Database(dbName).Collection(coll),
	}
}

func (r *repository) Insert(ctx context.Context, session *Session) (*Session, error) {
	result, err := r.collection.InsertOne(ctx, session)
	if err != nil {
		return nil, err
	}

	session.ID = result.InsertedID.(primitive.ObjectID)
	return session, nil
}
