package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Credential struct {
		Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		UserId string `json:"user_id" bson:"user_id"`
		RoleCode int `json:"role_code" bson:"role_code"`
		AccessToken string `json:"access_token" bson:"access_token"`
		RefreshToken string `json:"refresh_token" bson:"refresh_token"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Role struct {
		Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Title string `json:"title" bson:"title"`
		Code int `json:"code" bson:"code"`
	}

)
