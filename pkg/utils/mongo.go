package utils

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ConvertToObject is a function to convert string to ObjectID
func ConvertToObjectId(id string) primitive.ObjectID {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Failed to convert string '%s' to ObjectID: %v", id, err)
		return primitive.NilObjectID
	}
	return objectID
}
