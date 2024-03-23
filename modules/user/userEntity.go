package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Email string `json:"email" bson:"email"`
		Password string `json:"password" bson:"password"`
		Username string `json:"username" bson:"username"`
		CreateAt time.Time `json:"created_at" bson:"created_at"`
		UpdateAt time.Time `json:"updated_at" bson:"updated_at"`
		UserRole []UserRole `bson:"user_role"`
	}

	UserRole struct {
		RoleTitle string `json:"role_title" bson:"role_title"`
		RoleCode int `json:"role_code" bson:"role_code"`
	}

	UserProfileBson struct {
		Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Email string `json:"email" bson:"email"`
		Username string `json:"username" bson:"username"`
		CreateAt time.Time `json:"created_at" bson:"created_at"`
		UpdateAt time.Time `json:"updated_at" bson:"updated_at"`
	}

	UserSavingAccount struct {
		UserId string `json:"user_id" bson:"user_id"`
		Balance float64 `json:"balance" bson:"balance"`
	}
)