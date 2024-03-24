package main

import (
	"context"
	"log"
	"os"
	"tansan/config"
	"tansan/pkg/database/migration"
)

func main() {
	ctx := context.Background()

	// Initialize configuration
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Please provide the path to the .env file")
		}
		return os.Args[1]
	}())

	switch cfg.App.Name {
	case "user":
		migration.UserMigrate(ctx, &cfg)
	case "auth":
		migration.AuthMigrate(ctx, &cfg)
	case "book":
		migration.BookMigrate(ctx, &cfg)
	case "userbooks":
		migration.UserbooksMigrate(ctx, &cfg)
	case "payment":
		migration.PaymentMigrate(ctx, &cfg)
	}
}