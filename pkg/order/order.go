package order

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
}
