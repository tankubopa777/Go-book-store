package userbooks

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Userbooks struct {
		Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		UserId string `json:"user_id" bson:"user_id"`
		BookId string `json:"book_id" bson:"book_id"`
	}
)