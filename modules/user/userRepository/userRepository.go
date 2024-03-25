package userRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"tansan/modules/user"
	"tansan/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	UserRepositoryService interface {
		IsUniqueUser(pctx context.Context, email, username string) bool
		InsertOneUser(pctx context.Context, req *user.User) (primitive.ObjectID, error)
		FindOneUserProfile(pctx context.Context, userId string) (*user.UserProfileBson, error)
	}

	userRepository struct {
		db *mongo.Client
	}
)

func NewUserRepository(db *mongo.Client) UserRepositoryService {
	return &userRepository{db: db}
}

func (r *userRepository) userDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("user_db")
}


// IsUniqueUser checks if the user is unique
func (r *userRepository) IsUniqueUser(pctx context.Context, email, username string) bool{

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	user := new(user.User)
	if err := col.FindOne(
		ctx,
		bson.M{"$or": []bson.M{
			{"username": username},
			{"email": email},
		}},
	).Decode(user); err != nil {
		log.Printf("Error: IsUniqueUser: %v", err)
		return true
	}
	return false
}

// InsertOneUser inserts a user into the database
func (r *userRepository) InsertOneUser(pctx context.Context, req *user.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	userId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOneUser: %v", err.Error())
		return primitive.NilObjectID, errors.New("error: inserting user failed")
	}
	return userId.InsertedID.(primitive.ObjectID), nil
}

func (r *userRepository) FindOneUserProfile(pctx context.Context, userId string) (*user.UserProfileBson, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.userDbConn(ctx)
	col := db.Collection("users")

	result := new(user.UserProfileBson)

	if err := col.FindOne(
		ctx,
		bson.M{"_id": utils.ConvertToObjectId(userId)},
		options.FindOne().SetProjection(
			bson.M{
				"_id":        1,
				"email":      1,
				"username":   1,
				"created_at": 1,
				"updated_at": 1,
			},
		),
	).Decode(result); err != nil {
		log.Printf("Error: FindOneUserProfile: %s", err.Error())
		return nil, errors.New("error: user profile not found")
	}

	return result, nil
}