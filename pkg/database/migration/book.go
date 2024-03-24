package migration

import (
	"context"
	"log"
	"tansan/config"
	"tansan/modules/book"
	"tansan/pkg/database"
	"tansan/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func bookDbConn(pctx context.Context, cfg *config.Config) *mongo.Database{
	return database.DbConn(pctx, cfg).Database("book_db")
}

func BookMigrate(pctx context.Context, cfg *config.Config) {
	db := bookDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("books")

	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"title", 1}}},
	})

	for _, index := range indexs {
		log.Printf("Index: %s", index)
	}

	// roles data
	documents := func() []any {
		roles := []*book.Book{
			{
				Title: "12 Rules for Life: An Antidote to Chaos",
				Price: 1000,
				ImageUrl : "https://th-test-11.slatic.net/p/2c864769347a733d9c4163dbec42012a.jpg",
				UsageStatus: true,
				Damage: 100,
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Title: "The Subtle Art of Not Giving a F*ck: A Counterintuitive Approach to Living a Good Life",
				Price: 2000,
				ImageUrl : "https://th-test-11.slatic.net/p/2c864769347a733d9c4163dbec42012a.jpg",
				UsageStatus: true,
				Damage: 100,
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Title: "Atomic Habits: An Easy & Proven Way to Build Good Habits & Break Bad Ones",
				Price: 3000,
				ImageUrl : "https://th-test-11.slatic.net/p/2c864769347a733d9c4163dbec42012a.jpg",
				UsageStatus: true,
				Damage: 100,
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Title: "The Four Agreements: A Practical Guide to Personal Freedom",
				Price: 4000,
				ImageUrl : "https://th-test-11.slatic.net/p/2c864769347a733d9c4163dbec42012a.jpg",
				UsageStatus: true,
				Damage: 100,
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
		}

		docs := make([]any, 0)
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate book completed: ", results.InsertedIDs)
}