package auth

import (
	"tansan/modules/user"
	"time"
)

type (
	UserLoginReq struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,min=8,max=32"`
	}

	RefreshTokenReq struct {
		RefreshTokenReq string `json:"refresh_token" form:"refresh_token" validate:"required.max=500"`
	}

	InsertUserRole struct {
		UserId string `json:"user_id" validate:"required"`
		RoleCode []int `json:"role_id" validate:"required"`
	}

	ProfileIntercepter struct {
		*user.UserProfile
		Credential *Credential `json:"credential"`
	}

	CredentialRes struct {
		Id 		 string `json:"id" bson:"_id,omitempty"`
		UserId       string `json:"user_id" bson:"user_id"`
		RoleCode     int `json:"role_code" bson:"role_code"`
		AccessToken  string `json:"access"`
		CreatedAt	time.Time `json:"created_at"`
		UpdatedAt	time.Time `json:"updated_at"`
	}
)