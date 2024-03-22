package main

import (
	"context"
	"log"
	"os"
	"tansan/config"
	"tansan/pkg/database"
	"tansan/server"
)

func main() {
	ctx := context.Background()
	_ = ctx

	// Initialize configuration
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Please provide the path to the .env file")
		}
		return os.Args[1]
	}())

	// Database connection
	 db := database.DbConn(ctx, &cfg)	
	 defer db.Disconnect(ctx)
	 log.Println(db)

	// Start server
	server.Start(ctx, &cfg, db)
}