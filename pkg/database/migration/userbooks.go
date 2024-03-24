package migration

import (
	"context"
	"log"
	"tansan/config"
	"tansan/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func userbooksDbConn(pctx context.Context, cfg *config.Config) *mongo.Database{
	return database.DbConn(pctx, cfg).Database("userbooks_db")
}

func UserbooksMigrate(pctx context.Context, cfg *config.Config) {
	db := userbooksDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("users_userbooks")

	indexs, err := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"user_id", 1}, {"book_id", 1}}},
	})
	if err != nil {
		log.Fatalf("Failed to create indexes: %v", err)
	}
	

	for _, index := range indexs {
		log.Printf("Index: %s", index)
	}

	col = db.Collection("users_userbooks_queue")

	results, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate userbooks completed: ", results)
}