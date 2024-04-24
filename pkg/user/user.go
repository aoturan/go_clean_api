package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
}

type CreateUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserParams(params CreateUserParams) (*User, error) {
	return &User{
		Name:  params.Name,
		Email: params.Email,
	}, nil
}
