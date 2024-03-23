package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ConvertToObject is a function to convert string to ObjectID
func ConvertToObject(id string) primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return objectID
}

