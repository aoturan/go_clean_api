package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const coll = "users"

type Repository interface {
	Insert(context.Context, *User) (*User, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client, dbName string) Repository {
	return &repository{
		collection: client.Database(dbName).Collection(coll),
	}
}

func (r *repository) Insert(ctx context.Context, user *User) (*User, error) {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}
