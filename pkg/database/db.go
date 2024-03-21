package database

import (
	"context"
	"log"
	"tansan/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConn(pctx context.Context, cfg *config.Config) *mongo.Client {
	ctx, cancle := context.WithTimeout(pctx, 10*time.Second)
	defer cancle()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Url))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return client
}