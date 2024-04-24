package message

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const coll = "messages"

type Repository interface {
	Insert(context.Context, *Message) (*Message, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client, dbName string) Repository {
	return &repository{
		collection: client.Database(dbName).Collection(coll),
	}
}

func (r *repository) Insert(ctx context.Context, message *Message) (*Message, error) {
	result, err := r.collection.InsertOne(ctx, message)
	if err != nil {
		return nil, err
	}

	message.ID = result.InsertedID.(primitive.ObjectID)
	return message, nil
}
