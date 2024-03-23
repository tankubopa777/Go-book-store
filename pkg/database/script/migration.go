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
	_ = ctx

	// Initialize configuration
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Please provide the path to the .env file")
		}
		return os.Args[1]
	}())

	switch cfg.App.Name {
	case "user":
	case "auth":
		migration.AuthMigrate(ctx, &cfg)
	case "book":
	case "userbook":
	case "payment":
	}
}