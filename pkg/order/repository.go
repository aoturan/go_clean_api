package order

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const coll = "orders"

type Repository interface {
	Insert(context.Context, *Order) (*Order, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client, dbName string) Repository {
	return &repository{
		collection: client.Database(dbName).Collection(coll),
	}
}

func (r *repository) Insert(ctx context.Context, order *Order) (*Order, error) {
	result, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}

	order.ID = result.InsertedID.(primitive.ObjectID)
	return order, nil
}
